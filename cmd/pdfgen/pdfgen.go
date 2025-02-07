package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/yohanesmario/CV/cmd/pdfgen/data"
	"github.com/yohanesmario/CV/cmd/pdfgen/mdgen"
	"gopkg.in/yaml.v2"
)

// Define the structure of the config YAML content
type Config struct {
	OutputName          string   `yaml:"output-name"`
	AtsKeywordThreshold float64  `yaml:"ats-keyword-threshold"`
	AtsSections         []string `yaml:"ats-sections"`
	AtsKeywords         []string `yaml:"ats-keywords"`
}

func main() {
	// Read and parse cv.yaml
	cvFile, err := os.ReadFile("src/cv.yaml")
	if err != nil {
		log.Fatalf("error reading cv.yaml: %v", err)
	}
	var cv data.CV
	err = yaml.Unmarshal(cvFile, &cv)
	if err != nil {
		log.Fatalf("error parsing cv.yaml: %v", err)
	}

	// Read and parse config.yaml
	configFile, err := os.ReadFile("src/config.yaml")
	if err != nil {
		log.Fatalf("error reading config.yaml: %v", err)
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("error parsing config.yaml: %v", err)
	}

	// Ensure the gen directory exists
	genDir := "gen"
	err = os.MkdirAll(genDir, 0755)
	if err != nil {
		log.Fatalf("error creating gen directory: %v", err)
	}

	// Parse the HTML template
	templateFile := "cmd/pdfgen/template.go.html"
	fmt.Printf("▶ ⏳ Reading %s...\n", templateFile)
	htmlTemplateBuf, err := os.ReadFile(templateFile)
	if err != nil {
		log.Fatalf("error reading template file: %v", err)
	}
	htmlTemplate := string(htmlTemplateBuf)
	tmpl, err := template.New("cv-template").Funcs(template.FuncMap{
		"join":    strings.Join,
		"replace": strings.ReplaceAll,
		"lower":   strings.ToLower,
	}).Parse(htmlTemplate)
	if err != nil {
		log.Fatalf("error parsing HTML template: %v", err)
	}

	// Execute the template with the data
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, cv)
	if err != nil {
		log.Fatalf("error executing template: %v", err)
	}

	// Write the HTML content to a file for debugging
	htmlFile := filepath.Join(genDir, "output.out.html")
	err = os.WriteFile(htmlFile, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("error writing HTML file: %v", err)
	}
	fmt.Printf("HTML file successfully generated: %s\n", htmlFile)

	// Generate PDF from the HTML content
	pdfFile := filepath.Join(genDir, fmt.Sprintf("%s.pdf", config.OutputName))
	generatePDF(htmlFile, pdfFile)

	// Generate Markdown file
	mdgen.GenerateMarkdown(config.OutputName)
}

// Function to generate PDF from HTML content
func generatePDF(inputFileName string, outputFilename string) {
	command := fmt.Sprintf(
		`wkhtmltopdf \
			--debug-javascript \
			--window-status ready_to_print \
			--dpi 300 \
			--orientation Portrait \
			--page-size A4 \
			--margin-top 12mm \
			--margin-bottom 12mm \
			--margin-left 12mm \
			--margin-right 12mm \
			"%s" "%s"`,
		inputFileName,
		outputFilename,
	)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error running shell script: %v", err)
	}

	fmt.Printf("PDF successfully generated: %s\n", outputFilename)
}

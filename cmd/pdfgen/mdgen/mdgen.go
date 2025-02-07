package mdgen

import (
	"fmt"
	"log"
	"os"

	"github.com/yohanesmario/CV/cmd/pdfgen/data"
	"gopkg.in/yaml.v2"
)

func GenerateMarkdown(outputName string) {
	// Read the YAML file
	yamlFile, err := os.ReadFile("src/cv.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Unmarshal the YAML file into the CV struct
	var cv data.CV
	err = yaml.Unmarshal(yamlFile, &cv)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Generate Markdown content
	md := fmt.Sprintf("# %s\n\n", cv.FullName)
	for _, contact := range cv.ContactInfo {
		md += fmt.Sprintf("- **%s:** [%s](%s)\n", contact.Label, contact.Value, contact.URI)
	}
	md += "\n## Professional Summary\n\n"
	md += fmt.Sprintf("%s\n\n", cv.Summary)
	md += "## Experiences\n\n"
	for _, exp := range cv.Experiences {
		md += fmt.Sprintf("### %s • *%s*\n", exp.Title, exp.Company)
		md += fmt.Sprintf("*%s* • *%s - %s*\n\n", exp.Location, exp.Start, func() string {
			if exp.Current {
				return "Present"
			}
			return exp.End
		}())
		for _, desc := range exp.Description {
			md += fmt.Sprintf("- %s\n", desc)
		}
		md += "\n*Tech-stack:* "
		for i, tech := range exp.TechStack {
			if i > 0 {
				md += ", "
			}
			md += tech
		}
		md += "\n\n"
	}
	md += "## Education\n\n"
	for _, edu := range cv.Educations {
		md += fmt.Sprintf("### %s • *%s*\n", edu.Degree, edu.Institution)
		md += fmt.Sprintf("*%s* • *%s - %s*\n\n", edu.Location, edu.Start, edu.End)
	}

	// Write the Markdown content to a file
	err = os.WriteFile("gen/"+outputName+".md", []byte(md), 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("Markdown successfully generated: gen/" + outputName + ".md")
}

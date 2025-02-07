package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	// Read the configuration file
	configFile := "src/config.yaml"
	fmt.Printf("▶ ⏳ Reading %s...\n", configFile)
	configContent, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("▶ ❌ Failed to read %s: %v\n", configFile, err)
		os.Exit(1)
	}

	// Parse the configuration
	var config map[string]interface{}
	err = yaml.Unmarshal(configContent, &config)

	if err != nil {
		fmt.Printf("▶ ❌ Failed to parse %s: %v\n", configFile, err)
		os.Exit(1)
	}

	filePath := "gen/" + config["output-name"].(string) + ".md"
	keywords := extractKeywords(config)
	sections := extractSections(config)

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("▶ ❌ Failed to read %s: %v\n", filePath, err)
		os.Exit(1)
	}

	text := string(content)

	if !hasRequiredSections(text, sections) {
		fmt.Println("▶ ❌ CV is missing key sections required for ATS compliance. Please update your CV.")
		os.Exit(1)
	}

	keywordPercentage := calculateKeywordPercentage(text, keywords)
	keywordThreshold := config["ats-keyword-threshold"].(float64)

	fmt.Printf("▶ 🔵 Total keywords: %d\n", len(keywords))
	fmt.Printf("▶ 🔵 Keyword match percentage: %.2f%%\n", keywordPercentage*100)
	fmt.Printf("▶ 🔵 Keyword threshold: %.2f%%\n", keywordThreshold*100)

	if keywordPercentage < keywordThreshold {
		fmt.Println("▶ ❌ CV is missing important keywords related to the job description. Please update your CV.")
		os.Exit(1)
	}

	fmt.Println("▶ ✅ CV is ATS compliant.")
}

// Check if the CV contains all required sections
func hasRequiredSections(text string, sections []string) bool {
	result := true
	for _, section := range sections {
		sectionPattern := fmt.Sprintf(`(?i)#+.*\s*%s.*\s*\n`, regexp.QuoteMeta(section))
		match, err := regexp.MatchString(sectionPattern, text)
		if err != nil {
			fmt.Printf("▶ ❌ Failed to match section pattern: %v\n", err)
			os.Exit(1)
		}
		if !match {
			result = false
			fmt.Printf("▶ 🟡 Missing section: %s\n", section)
		}
	}
	return result
}

// Check if the CV contains important keywords
func calculateKeywordPercentage(text string, keywords []string) float64 {
	total := float64(len(keywords))
	counter := 0.0
	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(text), strings.ToLower(keyword)) {
			counter += 1.0
		} else {
			fmt.Printf("▶ 🟡 Missing keyword: %s\n", keyword)
		}
	}
	if total > 0.0 {
		return counter / total
	} else {
		return 0.0
	}
}

// Extract keywords from the configuration file
func extractKeywords(config map[string]interface{}) []string {
	srcKeywords := config["ats-keywords"].([]interface{})
	keywords := make([]string, len(srcKeywords))
	for i, keyword := range srcKeywords {
		keywords[i] = keyword.(string)
	}
	return keywords
}

// Extract sections from the configuration file
func extractSections(config map[string]interface{}) []string {
	srcSections := config["ats-sections"].([]interface{})
	sections := make([]string, len(srcSections))
	for i, section := range srcSections {
		sections[i] = section.(string)
	}
	return sections
}

#!/bin/bash

set -e

# Define styles
RESET='\033[0m'
GREEN='\033[1;32m'
CYAN='\033[1;36m'
YELLOW='\033[1;33m'

PROJECT_ROOT=$(git rev-parse --show-toplevel)

# Delete all generated files
echo -e "${CYAN}BUILD   ▶ Removing generated files...${RESET}"
find gen -name "*" -type f -delete

# Run PDF generation
echo -e "${CYAN}BUILD   ▶ Running PDF generation...${RESET}"
go run $PROJECT_ROOT/cmd/pdfgen/pdfgen.go

# Check for ATS compliance
echo -e "${CYAN}BUILD   ▶ Checking for ATS compliance...${RESET}"
go run $PROJECT_ROOT/cmd/atschecker/atschecker.go

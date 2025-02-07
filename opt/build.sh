#!/bin/bash

set -e

# Define styles
RESET='\033[0m'
GREEN='\033[1;32m'
CYAN='\033[1;36m'
YELLOW='\033[1;33m'

PROJECT_ROOT=$(git rev-parse --show-toplevel)

exit_on_no_change() {
    local changes_detected=false
    for file in "$@"; do
        if [[ -n $(git status -s | grep "$file") ]]; then
            changes_detected=true
            break
        fi
    done

    if ! $changes_detected; then
        echo -e "${GREEN}BUILD   ▶ No changes detected. Nothing to build.${RESET}"
        exit 0
    fi
}

# Check for git changes
echo -e "${CYAN}BUILD   ▶ Checking for relevant changes...${RESET}"
exit_on_no_change \
    "cv.yaml" \
    "config.yaml" \
    "pdfgen.go" \
    "mdgen.go" \
    "template.go.html" \
    "build.sh"
echo -e "${YELLOW}BUILD   ▶ Changes detected.${RESET}"
git status -s

# Delete all generated files
echo -e "${CYAN}BUILD   ▶ Removing generated files...${RESET}"
find gen -name "*" -type f -delete

# Run PDF generation
echo -e "${CYAN}BUILD   ▶ Running PDF generation...${RESET}"
go run $PROJECT_ROOT/cmd/pdfgen/pdfgen.go

# Check for ATS compliance
echo -e "${CYAN}BUILD   ▶ Checking for ATS compliance...${RESET}"
go run $PROJECT_ROOT/cmd/atschecker/atschecker.go

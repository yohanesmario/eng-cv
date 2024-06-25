#!/bin/bash

set -e

# Define styles
RESET='\033[0m'
CYAN='\033[1;36m'

# Commit changes
echo -e "${CYAN}COMMIT  ▶ Committing changes...${RESET}"
git add .
git commit-status

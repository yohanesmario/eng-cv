#!/bin/bash

set -e

# Define styles
RESET='\033[0m'
CYAN='\033[1;36m'

echo -e "${CYAN}PUSH    ▶ Force pushing to remote repository...${RESET}"
git push origin $(git rev-parse --abbrev-ref HEAD) --force

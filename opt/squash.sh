#!/bin/bash

set -e

# Define styles
RESET='\033[0m'
CYAN='\033[1;36m'

echo -e "${CYAN}SQUASH  ▶ Squashing commits...${RESET}"
git reset $(git commit-tree HEAD^{tree} -m "squashed")

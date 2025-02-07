#!/bin/bash

# Define styles
RESET='\033[0m'
RED='\033[1;31m'
GREEN='\033[1;32m'
CYAN='\033[1;36m'
YELLOW='\033[1;33m'

echo -e "${CYAN}INIT    ▶ Checking if .git/config has been initialized...${RESET}"
if grep -q "../.gitconfig" .git/config; then
    echo -e "${GREEN}INIT    ▶ .git/config has been initialized. Skipping initialization.${RESET}"
    exit 0
fi

echo -e "${CYAN}INIT    ▶ Installing local git config...${RESET}"
git config --local include.path ../.gitconfig
echo -e "${YELLOW}INIT    ▶ Updated local git config: ${RESET}"
cat .git/config

echo -e "${GREEN}INIT    ▶ Local git config installed.${RESET}"

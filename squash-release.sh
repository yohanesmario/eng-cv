#!/bin/bash

# Define styles
RESET='\033[0m'
RED='\033[1;31m'
GREEN='\033[1;32m'
CYAN='\033[1;36m'
YELLOW='\033[1;33m'

# Function to handle errors
handle_error() {
    echo -e "${RED}RELEASE ▶ An error occurred. Changes not pushed to remote repository.${RESET}"
    exit 1
}

# Trap errors and call handle_error
trap 'handle_error' ERR

# Making sure init.sh has been executed
./opt/init.sh

# Check for git changes
echo -e "${CYAN}RELEASE ▶ Checking for changes...${RESET}"
if [[ -z $(git status -s) ]]; then
    echo -e "${GREEN}RELEASE ▶ No changes detected. Nothing to push.${RESET}"
    exit 0
else
    echo -e "${YELLOW}RELEASE ▶ Changes detected.${RESET}"
    git status -s
fi

# Build
./opt/build.sh

# Commit changes
./opt/commit.sh

# Squash commits
./opt/squash.sh

# Push changes to remote repository
./opt/force-push.sh

# Success message
echo -e "${GREEN}RELEASE ▶ Squashed changes pushed to remote repository.${RESET}"

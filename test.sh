#!/bin/bash

# Function to ask for confirmation
confirm_action() {
    while true; do
        read -p "Do you want to proceed to the next step? (yes/no): " yn
        case $yn in
            [Yy]* ) return 0;; # Success/Proceed
            [Nn]* ) return 1;; # Failure/Stop
            * ) echo "Please answer yes or no.";;
        esac
    done
}

# Implementation
echo "Starting the process..."

if confirm_action; then
    echo "Action confirmed! Moving to the next step..."
    # Add your next steps here
else
    echo "Action cancelled by user."
    exit 1
fi
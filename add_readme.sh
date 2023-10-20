#!/bin/bash

# Define the filename for the README file
readme_file="README.md"

# Find all directories and subdirectories, excluding the current directory ('.') to avoid overwriting the README in the current directory
find . -type d ! -path . | while read -r dir; do
  # Create the README file in each directory
  touch "$dir/$readme_file"
  # Optionally, add content to the README file
  echo "# ${dir:2}" > "$dir/$readme_file"
done

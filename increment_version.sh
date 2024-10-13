#!/bin/bash

# Read the current version
current_version=$(cat version.txt)

# Extract the version number and increment it
version_number=$(echo $current_version | sed 's/v//')  # Remove 'v'
new_version_number=$((version_number + 1))            # Increment
new_version="v$new_version_number"                      # Add 'v' back

# Write the new version to the version.txt file
echo $new_version > version.txt

echo $new_version  # Print the new version for use in the build

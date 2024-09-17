#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Navigate to the frontend directory
cd frontend || exit 1

# Run the build command using bun
bun run build

# Navigate to the backend/public directory
cd ../backend/public || exit 1

# Delete index.html
rm -f index.html

# Go back to the root directory
cd ../../..

# Move mango.exe from backend to the root directory
mv backend/mango.exe .

echo "Script executed successfully."

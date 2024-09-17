#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

echo "Starting build script..."

# Navigate to the frontend directory
cd frontend || exit 1
echo "In frontend directory: $(pwd)"

# Run the build command using bun
bun run build

# Navigate to the backend directory
cd ../backend || exit 1
echo "In backend directory: $(pwd)"

# Build the Go backend
go build -o mango

# Navigate to the public directory
cd public || exit 1
echo "In public directory: $(pwd)"

# Delete index.html
rm -f index.html

# Go back to the root directory
cd ../..
echo "Back to root directory: $(pwd)"

# Move mango executable from backend to the root directory
mv backend/mango .

echo "Script executed successfully."
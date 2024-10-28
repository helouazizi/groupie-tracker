#!/bin/bash

# Find and kill the running Go server
pkill -f "go run ."  # Adjust if you need a more specific match

# Wait for a moment to ensure the process is terminated
#sleep 1

# Run the Go server again
#go run .
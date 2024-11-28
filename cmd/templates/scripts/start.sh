#!/bin/bash

# Default values
CONFIG_FILE="config/config.yaml"
LOG_LEVEL="info"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        -c|--config)
            CONFIG_FILE="$2"
            shift
            shift
            ;;
        -l|--log-level)
            LOG_LEVEL="$2"
            shift
            shift
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Start the server
./bin/server --config "$CONFIG_FILE" --log-level "$LOG_LEVEL" 
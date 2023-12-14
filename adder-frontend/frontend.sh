#!/bin/bash

# Define color variables
GREEN='\033[0;32m'
AMBER='\033[0;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Function to print usage
print_usage() {
    echo -e "${AMBER}Usage: $0 [DOCKER_ARGS...] [ADDER_ARGS]${NC}"
    echo -e "${AMBER}If no arguments or 'default' is provided, a default Docker run command will be used.${NC}"
}

# Default Docker arguments
md5time=$(date +%s%N | xxd -p | tail -c 7)
DEFAULT_DOCKER_ARGS="--name adder-frontend-${md5time} -p 8500:3000 -e HOST_ADDRESS=$HOST"

# Default ADDER argument
DEFAULT_ADDER_ARGS=" "

# Check if no arguments or 'default' is provided
if [ $# -eq 0 ] || [[ "$1" =~ ^[dD][eE][fF][aA][uU][lL][tT]$ ]]; then
    DOCKER_ARGS=$DEFAULT_DOCKER_ARGS
    ADDER_ARGS=$DEFAULT_ADDER_ARGS
else
    # Extract ADDER_ARGS (last argument)
    ADDER_ARGS="${@: -1}"
    # Remove the last argument to get DOCKER_ARGS
    DOCKER_ARGS="${@:1:$#-1}"
fi
echo hi
# Run Docker command
echo -e "${GREEN}Running Docker container...${NC}"
command="docker run ${DOCKER_ARGS} adder-frontend ${ADDER_ARGS}"

echo -e "${AMBER}Executing command: ${command}${NC}"

# Execute the command
eval "$command &"


if [ $? -eq 0 ]; then
    echo -e "${GREEN}Docker container started successfully.${NC}"
else
    echo -e "${RED}Failed to start Docker container.${NC}"
    exit 1
fi

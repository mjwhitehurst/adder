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
#take a hash of current time so instances have unique names - doesn't have to be secure, just unique enough
md5time=$(date +%s%N | xxd -p | tail -c 7)
DEFAULT_DOCKER_ARGS="--name adder-backend-${md5time} --cap-add=SYS_PTRACE --user \"$(id -u):$(id -u)\" -p 8080:8080 -p 2345:2345 -v $SRC:/app/sourcedir"

# Default ADDER argument
DEFAULT_ADDER_ARGS="dlv debug --headless --listen=:2345 --api-version=2 --log"

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

# Run Docker command
echo -e "${GREEN}Running Docker container...${NC}"
command="docker run ${DOCKER_ARGS} adder-backend ${ADDER_ARGS}"

echo -e "${AMBER}Executing command: ${command}${NC}"

# Execute the command
eval "$command &"


if [ $? -eq 0 ]; then
    echo -e "${GREEN}Docker container started successfully.${NC}"
else
    echo -e "${RED}Failed to start Docker container.${NC}"
fi

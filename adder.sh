#!/bin/bash

# Define color variables
GREEN='\033[0;32m'
AMBER='\033[0;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

### FUNCTIONS ###

# Function to kill a Docker container by image name
kill_container() {
    local image_name=$1
    echo -e "Searching for containers running image: ${image_name}"

    # Find container IDs for given image name
    container_ids=$(docker ps -q --filter ancestor=${image_name})

    # Check if any containers were found
    if [ -z "$container_ids" ]; then
        echo -e "${GREEN}No running containers found for image: ${image_name}${NC}"
    else
        # Stop the containers
        echo "Stopping containers: $container_ids"
        docker stop $container_ids

        if [ $? -eq 0 ]; then
            echo -e "${GREEN}Successfully stopped container(s) for image: ${image_name}${NC}"
        else
            echo -e "${RED}Failed to stop container(s) for image: ${image_name}${NC}"
        fi
    fi
}


### START OF ACTUAL SCRIPT ###

# Check current directory and get adder directory
current_dir=$(basename "$PWD")
base_dir=$(basename "$current_path")
if [[ "$base_dir" == "adder" ]]; then
    adder_path="$current_path"
else
    adder_path=$(dirname "$current_path")
fi


if [[ "$current_dir" != "adder" && "$current_dir" != "adder-backend" && "$current_dir" != "adder-frontend" ]]; then
    echo -e "${RED}Error: This script must be run from a directory named adder, adder-backend, or adder-frontend.${NC}"
    exit 1
fi

# Check if at least one argument is provided
if [ $# -lt 1 ]; then
    echo -e "${AMBER}Usage: $0 [backend|frontend|start] [...]${NC}"
    exit 1
fi

# Handle the first argument
case "$1" in
    backend)
        shift # Remove the first argument
        # Run backend.sh with remaining arguments
        eval ". ${adder_path}/adder-backend/backend.sh $@ &"
        exit 0
        ;;
    frontend)
        shift # Remove the first argument
        # Run frontend.sh with remaining arguments
        eval ". ${adder_path}/adder-frontend/frontend.sh $@ &"
        exit 0

        ;;
    start)
        # Run backend.sh and frontend.sh with 'default' argument
        echo -e "${GREEN} starting frontend"
        eval ". ${adder_path}/adder-backend/backend.sh default"
        echo -e "${GREEN} starting backend"
        eval ". ${adder_path}/adder-frontend/frontend.sh default"
        exit 0

        ;;
    kill)
        # Kill both frontend and backend processes
        echo -e "${RED} killing backend"
        kill_container "adder-backend"
        echo -e "${RED} killing fronend"
        kill_container "adder-frontend"
        echo -e "${NC}"
        exit 0
        ;;
    build)
        echo -e "${AMBER} - Rebuilding adder script... ${NC}"
        eval ". build_adder_script.sh"
        echo -e "${GREEN}done"
        echo -e "${AMBER} - building backend"
        eval "docker build -t adder-backend ${adder_path}/adder-backend"
        echo -e "${GREEN}done"
        echo -e "${AMBER} - building frontend"
        eval "docker build -t adder-frontend ${adder_path}/adder-backend"
        echo -e "${GREEN}done"
        ;;
    *)
        echo -e "${RED}Invalid argument: $1${NC}"
        echo -e "${AMBER}Usage: $0 [backend|frontend|start] [...]${NC}"
        exit 1
        ;;
esac

echo -e "${GREEN}Adder Script finished${NC}"



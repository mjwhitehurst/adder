#!/bin/bash


# Define color variables
GREEN='\033[0;32m'
AMBER='\033[0;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

### FUNCTIONS ###

# Coloured text
function success( ){
    printf "${GREEN}$1${NC}\n"
}

function warn() {
    printf "${AMBER}$1${NC}\n"
}

function err() {
    printf "${RED}$1${NC}\n"
}

## not a function, but 'overrides' for exit and return in diff situations.
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
        exit_adder='exit'
    else
        exit_adder='return'
fi

# Function to kill a Docker container by image name
function kill_container() {
    local image_name=$1
    warn "Searching for containers running image: ${image_name}"

    # Find container IDs for given image name
    container_ids=$(docker ps -q --filter ancestor=${image_name})

    # Check if any containers were found
    if [ -z "$container_ids" ]; then
        success "No running containers found for image: ${image_name}"
    else
        # Stop the containers
        echo "Stopping containers: $container_ids"
        docker stop $container_ids

        if [ $? -eq 0 ]; then
            success "Successfully stopped container(s) for image: ${image_name}"
        else
            err "Failed to stop container(s) for image: ${image_name}"
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
    err "Error: This script must be run from a directory named adder, adder-backend, or adder-frontend."
    $exit_adder 1
fi

# Check if at least one argument is provided
if [ $# -lt 1 ]; then
    warn "Usage: $0 [backend|frontend|start|kill] [...]"
    $exit_adder 1
fi

# Handle the first argument
case "$1" in
    backend)
        shift # Remove the first argument
        # Run backend.sh with remaining arguments
        eval ". ${adder_path}/adder-backend/backend.sh $@ &"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
        $exit_adder 0
        ;;
    frontend)
        shift # Remove the first argument
        # Run frontend.sh with remaining arguments
        eval ". ${adder_path}/adder-frontend/frontend.sh $@ &"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
        $exit_adder 0

        ;;
    start)
        # Run backend.sh and frontend.sh
        success "starting frontend"
        eval "adder frontend"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1

        success "starting backend"
        eval "adder backend"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
        $exit_adder 0

        ;;
    kill)
        # Kill both frontend and backend processes
        err " killing backend"
        kill_container "adder-backend"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
        err " killing fronend"
        kill_container "adder-frontend"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
        $exit_adder 0
        ;;
    build)
        # When we run build, we intend to come here twice.
        #  once with the 'current' stuff to move the 'new' code into $BIN
        if [ "$2" != "-skip_script_build" ]; then
            warn " - Rebuilding adder script..."
            eval ". build_adder_script.sh"
            [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
            success "done"
            #then the second time, so we are running the 'new' code when building.
            eval "adder build -skip_script_build" # <-- so this...
            $exit_adder $?
        fi
        # will take us here.
        warn " - building backend NEW"
        eval "docker build -t adder-backend ${adder_path}/adder-backend"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
        success "done"
        warn " - building frontend"
        eval "docker build -t adder-frontend ${adder_path}/adder-frontend"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
        success "done"
        ;;
    cd)
        success " Adder ⤵⤵⤵"
        cd ~/adder
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1        $exit_adder 0
        ;;
    *)
        err "Invalid argument: $1"
        warn "Usage: $0 [backend|frontend|start] [...]"
        ;;
esac

success "Adder Script finished"
$exit_adder 0

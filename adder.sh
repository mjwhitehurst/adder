#!/bin/bash

BASHRC_CHANGED="FALSE"

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

add_to_bashrc() {
    local string_to_add="$2 #$1"
    local string_tag="$1"
    local bashrc="${HOME}/.bashrc"
    local adder_marker="##>>ADDER<<##"

    # Check if the string already exists in .bashrc
    if grep -qF -- "$string_tag" "$bashrc"; then
        return 1
    fi

    # Check if the ADDER marker exists
    if grep -qF -- "$adder_marker" "$bashrc"; then
        # Use awk to insert the line on the next free line after the marker
        awk -v line="$string_to_add" -v marker="$adder_marker" '
            $0 ~ marker && !added {
                print; getline; print line; added=1; next
            }
            {print}
        ' "$bashrc" > tmpfile && mv tmpfile "$bashrc"
    else
        # Append the line to the end of the file
        echo "$string_to_add" >> "$bashrc"
    fi

    BASHRC_CHANGED="TRUE"
    success "$1 added to .bashrc"
}

## Function for adding default adder stuff to bashrc, assuming it exists.
function set_up_bashrc(){
    #$q is a quote in .bashrc.
    local q='\"'
    add_to_bashrc "#>>ADDER<<##"
    #assume we're not going to add more than 6 args...
    add_to_bashrc "adder_cf" "alias adder='f(){ if [ ${q}\$1${q} = ${q}cd${q} ]; then cd \$HOME/adder; else adder ${q}\$1${q} ${q}\$2${q} ${q}\$3${q} ${q}\$4${q} ${q}\$5${q} ${q}\$6${q}; fi; }; f'"
}

function printbuildhelp(){
    warn " -- BUILDING -- "
    warn ""
    warn " Can be run with no arguments to 'just do' everything "
    warn "  or with the following args:"
    warn ""
    warn "  * -so   / -scrip_tonly         - move this script into \$BIN to call with 'adder'"
    warn "  * -ssb  / -skip_script_build   - ignore this script and move on"
    warn ""
    warn " --------------"
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

#checking args for inputs:
function check_valid_args() {
    local valid_args=("$@")  # Capture all arguments
    local arg_count=$#  # Number of arguments passed to the function

    # Find the index of '--' which separates valid arguments from script arguments
    local separator_index=0
    for (( i=1; i<=arg_count; i++ )); do
        if [ "${valid_args[$i-1]}" = "--" ]; then
            separator_index=$i
            break
        fi
    done

    # If '--' not found, return error
    if [ $separator_index -eq 0 ]; then
        echo "Error: Arguments separator '--' not found."
        return 1
    fi

    # Split arguments into two arrays: valid arguments and script arguments
    local script_args=("${valid_args[@]:$separator_index}")
    valid_args=("${valid_args[@]:0:$((separator_index - 1))}")

    # Remove empty arguments from script_args
    script_args=($(echo "${script_args[@]}" | tr ' ' '\n' | grep -v "^$"))

    local all_valid=true
    for arg in "${script_args[@]}"; do
        local found=false
        for valid_arg in "${valid_args[@]}"; do
            if [ "$arg" == "$valid_arg" ]; then
                found=true
                break
            fi
        done

        if [ "$found" = false ]; then
            all_valid=false
            echo "Invalid argument detected: $arg"
            break
        fi
    done

    if [ "$all_valid" = true ]; then
        return 0
    else
        return 1
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

#Aliases etc.
set_up_bashrc

if [[ "$current_dir" != "adder" && "$current_dir" != "adder-backend" && "$current_dir" != "adder-frontend" && "$1" != "cd" ]]; then
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
        #check args
        if [ $# -gt 0 ]; then #if we're given arguments
            #check they're accepted:
            valid_args=(    "build"
                            "-skip_script_build" "-ssb"
                            "-script_only"       "-so" )

            check_valid_args "${valid_args[@]}" -- "$@"

            [ $? == 1 ] && printbuildhelp && $exit_adder 1
        fi

        # When we run build, we intend to come here twice.
        #  once with the 'current' stuff to move the 'new' code into $BIN
        if ! [[ " $* " =~ " -ssb " ]] && ! [[ " $* " =~ " -skip_script_build " ]]; then
            warn " - Rebuilding adder script..."
            eval ". build_adder_script.sh"
            [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
            success "done"
            #then the second time, so we are running the 'new' code when building.
            eval "adder build -skip_script_build $@" # <-- so this...
            $exit_adder $?
        fi
        # will take us here.


        if [[ " $* " =~ " -so " ]] || [[ " $* " =~ " -script_only " ]]; then
            $exit_adder 0
        fi

        warn " - building backend"
        #eval "docker build -t adder-backend ${adder_path}/adder-backend"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
        success "done"
        warn " - building frontend"
        #eval "docker build -t adder-frontend ${adder_path}/adder-frontend"
        [ $? == 1 ] && err ">FAILED<" && $exit_adder 1
        success "done"
        ;;
    *)
        err "Invalid argument: $1"
        warn "Usage: $0 [backend|frontend|start] [...]"
        ;;
esac

[ "${BASHRC_CHANGED}" == "TRUE" ] && err "BASHRC CHANGED" && warn "re-run \". ~/.bashrc\" or log out and back in." && $exit_adder 1


success "Adder Script finished"
$exit_adder 0

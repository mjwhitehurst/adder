## USAGE: . build_adder_script.sh ##
## WHY? so you can just call adder from wherever.

# Define color variables
GREEN='\033[0;32m'
AMBER='\033[0;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

## FUNCTIONS ##
add_to_path() {
    echo "export PATH=\"\$PATH:$1\"" >> ~/.bashrc
    echo -e "${GREEN}Added $1 to PATH in .bashrc.${NC}"
}

create_directory() {
    mkdir -p "$1"
    echo -e "${GREEN}Created directory $1.${NC}"
}

##START OF MAIN SCRIPT ##


# Check if the current directory is 'adder'
current_dir=$(basename "$PWD")
if [ "$current_dir" != "adder" ]; then
    echo -e "${RED}This script must be run from the 'adder' directory. Exiting.${NC}"
fi

# Check if $BIN is set and if it exists
if [ -z "$BIN" ]; then
    echo -e "${RED}\$BIN is not set. Exiting.${NC}"
fi

# Check if $BIN exists
if [ -d "$BIN" ]; then
    # Check if $BIN is in PATH
    if [[ ":$PATH:" != *":$BIN:"* ]]; then
        echo -e "${AMBER}Warning: Directory specified in \$BIN exists but is not in PATH.${NC}"
        read -rp "Add $BIN to PATH in .bashrc? (Y/N): " choice
        if [[ "$choice" =~ ^[Yy]$ ]]; then
            add_to_path "$BIN"
        else
            echo -e "${RED}Not adding to PATH. Exiting.${NC}"
        fi
    fi
else
    # $BIN does not exist
    if [[ ":$PATH:" != *":$BIN:"* ]]; then
        echo -e "${RED}\$BIN does not exist and is not in PATH.${NC}"
        read -rp "Create $BIN and add to PATH in .bashrc? (Y/N): " choice
        if [[ "$choice" =~ ^[Yy]$ ]]; then
            create_directory "$BIN"
            add_to_path "$BIN"
        else
            echo -e "${RED}Not creating directory or adding to PATH. Exiting.${NC}"
        fi
    else
        echo -e "${AMBER}Warning: \$BIN does not exist but is in PATH.${NC}"
        read -rp "Create $BIN? (Y/N): " choice
        if [[ "$choice" =~ ^[Yy]$ ]]; then
            create_directory "$BIN"
        else
            echo -e "${RED}Not creating directory. Exiting.${NC}"
        fi
    fi
fi
# Path to the adder.sh script
adder_script="./adder.sh"

adder_backend="./adder-backend/backend.sh"
adder_frontend="./adder-frontend/frontend.sh"


# Move to $BIN
cp "$adder_script" "$BIN/adder"
chmod a+x $BIN/adder


if [ $? -eq 0 ]; then
    echo -e "${GREEN}adder moved to $BIN${NC}"
else
    echo -e "${RED}Failed to move adder.sh${NC}"
fi

cp "$adder_frontend" "$BIN/adder_frontend"
chmod a+x $BIN/adder_frontend


if [ $? -eq 0 ]; then
    echo -e "${GREEN}adder_frontend moved to $BIN${NC}"
else
    echo -e "${RED}Failed to move adder_frontend.sh${NC}"
fi

cp "$adder_backend" "$BIN/adder_backend"
chmod a+x $BIN/adder_backend


if [ $? -eq 0 ]; then
    echo -e "${GREEN}adder_backend moved to $BIN${NC}"
else
    echo -e "${RED}Failed to move adder_frontend.sh${NC}"
fi

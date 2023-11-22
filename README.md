# Adder
A Command Line + Server tool to automate simple code changes on a Matflo System
Install instructions for Linux (CentOS) only. Will update with other OS installs as we do them.

## System Requirements

* Docker Engine >= 24.0.4
* OS able to install go >= 1.16.4
* Existing Directories: ~/source
* Source directory containing Matflo files - currently only xxx_definitions.h files edited




# Installation

## Go - Not necessary, but useful.
Installed go by downloading:

'''bash
    sudo yum update
    wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.16.4.linux-amd64.tar.gz


Added to .bashrc:

'''bash
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin

Run .bashrc:

'''bash
    . .bashrc

Check install:

'''bash
    go version

## Docker
Install docker


add user to docker users:
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
docker run hello-world

build:

move to adder-backend directory
touch go.sum
docker build -t adder-backend .

move to adder-frontend directory



run (cmd line):
docker run                            \
    --cap-add=SYS_PTRACE              \
    --user "$(id -u)":"$(id -u)"      \
    -v $SRC:/app/sourcedir            \
     adder-backend                    \
    GET_ALL_FIELDS TM

run (server):
docker run                        \
    --cap-add=SYS_PTRACE          \
    --user "$(id -u)":"$(id -u)"  \
    -p 8080:8080                  \
    -p 2345:2345                  \
    -v $SRC:/app/sourcedir        \
    adder-backend                 \
    dlv debug --headless --listen=:2345 --api-version=2 --log

curls (server):
curl -X POST http://localhost:8080/add-db-field \
     -H 'Content-Type: application/json'        \
     -d '{"database_name":"tm", "field_name":"MyBool", "field_type":"int", "comment":"TESTCOMMENT", "option":"NONDB"}'

run frontend:

docker run              \
  -p 8500:3000          \
  -e HOST_ADDRESS=$HOST \
  adder-frontend




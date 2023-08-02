# adder
to automate simple code changes on a system with certain system layouts.



# this is just my notepad of what to do. think I'm going to put go install into the docker image so can run anywhere docker can
installed go by downloading:

sudo yum update
wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.16.4.linux-amd64.tar.gz

added to .bashrc:
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

run .bashrc, then run
go version



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



run (cmd line) (TODO - UPDATE THIS WITH THINGS FROM SERVER):
docker run --user "$(id -u)":"$(id -u)" -v $SRC:/app/sourcedir adder-backend ADD_REC_FIELD TM Flag1 BOOLEAN hi

run (server):

docker run                \
    --cap-add=SYS_PTRACE  \
    -e HOST_UID=$(id -u)  \
    -e HOST_GID=$(id -g)  \
    -p 8080:8080          \
    -p 2345:2345          \
    -v $SRC:/app/sourcedir\
    adder-backend         \
    dlv debug --headless --listen=:2345 --api-version=2 --log

curls (server):
curl -X POST http://localhost:8080/add-db-field -H 'Content-Type: application/json' -d '{"database_name":"tm", "field_name":"MyBool", "field_type":"int", "comment":"TESTCOMMENT", "option":"NONDB"}'

run frontend:

docker run              \
  -p 8500:3000          \
  -e HOST_ADDRESS=$HOST \
  adder-frontend




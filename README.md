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



run:
docker run --user "$(id -u)":"$(id -u)" -v $SRC:/app/sourcedir adder ADD_REC_FIELD TM Flag1 BOOLEAN hi





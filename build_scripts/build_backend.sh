#!/bin/bash

export GOPATH=$PWD
go get fox/main
go build -i -o bin/fox fox/main
go get login/main
go build -i -o bin/login login/main

mkdir -p bin/config/$USER
cp src/config/config.json.template bin/config/$USER/config.json
mkdir /tmp/foxdb
go run src/authn/keygen/KeyGen.go > key.base64

cd bin
./fox
./login
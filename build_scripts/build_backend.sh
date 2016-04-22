#!/bin/bash

export GOPATH=$PWD

go get fox/foxservice
go get login/loginservice

mkdir -p bin/config/$USER
cp src/config/config.json.template bin/config/$USER/config.json

mkdir /tmp/foxdb  
./bin/foxservice

go run src/authn/keygen/KeyGen.go > key.base64 
./bin/loginservice

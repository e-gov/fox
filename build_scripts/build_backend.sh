#!/bin/bash

export GOPATH=$PWD

go get fox/foxservice
go get login/loginservice

go get -d -t fox/fox_test
go get -d -t login/login_test

mkdir -p bin/config/$USER
cp src/config/config.json.template bin/config/$USER/config.json
cp src/config/config.json.template bin/config/$USER/test_config.json

mkdir /tmp/foxdb  
./bin/foxservice

go run src/authn/keygen/KeyGen.go > key.base64 
./bin/loginservice

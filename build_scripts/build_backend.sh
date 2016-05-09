#!/bin/bash

export GOPATH=$PWD

go get fox/foxservice
go get login/loginservice

go get -d -t fox/fox_test
go get -d -t login/login_test

mkdir -p config/$USER
cp -n src/config/config.json.template config/$USER/config.json
cp -n src/config/config.json.template config/$USER/test_config.json

mkdir /tmp/foxdb  
./bin/foxservice

go run src/authn/keygen/KeyGen.go > key.base64 
./bin/loginservice

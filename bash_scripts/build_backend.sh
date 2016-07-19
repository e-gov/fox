#!/bin/bash

export GOPATH=$PWD

go get github.com/e-gov/fox/fox/foxservice
go get github.com/e-gov/fox/login/loginservice

go get -d -t github.com/e-gov/fox/fox/fox_test
go get -d -t github.com/e-gov/fox/login/login_test

mkdir -p config/$USER
cp -n config/config.json.template config/$USER/config.json
cp -n config/config.json.template config/$USER/test_config.json

mkdir src/github.com/e-gov/fox/tmp/foxdb  
./bin/foxservice

go run src/github.com/e-gov/fox/authn/keygen/KeyGen.go > key.base64 
./bin/loginservice

export GOPATH=$PWD
go get fox/main
go build -i -o bin/fox.exe fox/main
go get login/main
go build -i -o bin/login.exe login/main

mkdir -p bin/config/$env:username
cp src/config/config.json.template bin/config/$env:username/config.json
mkdir /tmp/foxdb
go run src/authn/keygen/KeyGen.go > key.base64

cd bin
start fox.exe
start login.exe
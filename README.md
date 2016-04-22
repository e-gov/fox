# Fox
An implementation of Fox Registry, the reference architecture for cloud-ready government systems. More detail of the project, its architecture and rationale is available at https://www.ria.ee/riigiarhitektuur/wiki/doku.php?id=an:rebasteregister [![Build Status](https://travis-ci.org/e-gov/fox.svg?branch=master)](https://travis-ci.org/e-gov/fox)

The FoxAPI application implements [this specification](http://editor.swagger.io/#/?import=https:%2F%2Fraw.githubusercontent.com%2Fe-gov%2Ffox%2Fmaster%2Ftatic%2F_data%2FFoxAPI.yaml)

There are two key components:
 1. The FoxService that implements the main business logic of the fox registry
 2. The LoginService that mints tokens for FoxService instances to be used and communicates with external authentication providers

## Starting a web-based UI

A web UI is built with `grunt`, to start it:

1. Make sure node.js and npm is installed.
2. Go to static folder `cd static`.
3. Run following commands

```
npm install
npm install -g grunt-cli
```
This should install all required tools to start the UI
Now go and copy `static/properties.json.sample` to `static/properties.json`
If you have default settings then sample properties will do just fine. Otherwise dig into properties.json

Now you should be all set. 
To run web ui run following command in `/static`
```
grunt serve
```
It should run webserver in `localhost:9000`



If you see errors about the encoding of files on OS X, try this:

```
export LC_ALL=en_US.UTF-8
export LANG=en_US.UTF-8
```

## Building a demo REST server

1. Change to the directory where the repository is cloned.
2. Setup environment and build application:

```
export GOPATH=$PWD
go get fox/foxservice
go get login/loginservice
```

## Running a REST server

1. Create folder config/{username} in bin, then copy and adapt example configuration file. 
2. Execute Fox binary passing an instance name as a parameter.

Linux
```
mkdir -p bin/config/$USER
cp src/config/config.json.template bin/config/$USER/config.json
```
Windows
```
mkdir -p bin/config/$env:username 
cp src/config/config.json.template bin/config/$env:username/config.json 
```

```
mkdir /tmp/foxdb  # make sure that the configured storage folder exists.
./bin/foxservice

go run src/authn/keygen/KeyGen.go > key.base64 # Generate the keyfile for authentication tokens
./bin/loginservice
```

REST interface will respond on **http://localhost:8090/**. You should now be able to use web UI in **http://localhost:9000/**.
To change a port or name of the application ("my" by default), check **./bin/foxservice -h**.

## Configuration
Configuration is user-based, every user has a folder with their username under config/, where their personal config file(s) live.
All services and tests use the same configuration file: config.{ext} for services, test_config.{ext} for tests. Config files can be in all formats supported by [Viper](http://github.com/spf13/viper) (JSON, TOML, YAML, HCL, Java properties).

## Reloading configuration
To reload configuration, both the login and fox services accept a HUP signal that should have both produce log messages about re-loading configuration

## Generating tokens for backend use
To generate tokens for headless clients, use  TokenMint.go:
```
go run src/authn/mint/TokenMint.go -key <a file containing a minting key> -user <a username the token should be assigned to>
```

## Generating passwords for authentication
To use the basic password authentication provider, passwords must be hashed and strored on server side. This happens like so:

```
touch pwd.list
go run src/authn/pwd/main/pwdMaker.go -user <username> -pwd <password> >> pwd.list
```

The pwd.list is a file referred to by the authn.PwdProvider.PwdFileName key in the config

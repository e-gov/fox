# Fox
An implementation of Fox Registry, the reference architecture for cloud-ready government systems. More detail of the project, its architecture and rationale is available at https://www.ria.ee/riigiarhitektuur/wiki/doku.php?id=an:rebasteregister [![Build Status](https://travis-ci.org/e-gov/fox.svg?branch=master)](https://travis-ci.org/e-gov/fox)

The FoxAPI application implements [this specification](http://editor.swagger.io/#/?import=https:%2F%2Fraw.githubusercontent.com%2Fe-gov%2Ffox%2Fmaster%2Ftatic%2F_data%2FFoxAPI.yaml)

There are two key components:
 1. The FoxService that implements the main business logic of the fox registry
 2. The LoginService that mints tokens for FoxService instances to be used and communicates with external authentication providers

## Starting a web-based UI

A web UI is built with jekyll, to start it:

1. Make sure Jekyll is installed.
2. Run `jekyll serve` from the static folder.

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
go get fox/main
go build -i -o fox fox/main  # the fox service will be built into ./fox

go get login/main
go build -i -o login login/main # the login service will be built into ./login
```

## Running a REST server

1. Copy and adapt example configuration file:
2. Execute Fox binary passing an instance name as a parameter.

```
cp src/fox/config.gcfg.template bin/config.gcfg
mkdir /tmp/foxdb  # make sure that the configured storage folder exists.
./fox


go run src/authn/keygen/KeyGen.go > key.base64 # Generate the keyfile for authentication tokens
./login
```

REST interface will respond on **http://localhost:8090/**. You should now be able to use web UI.
To change a port or name of the application ("my" by default), check **./bin/fox -h**.

## Reloading configuration
To reload configuration, both the login and fox services accept a HUP signal that should have both produce log messages about re-loading configuration

## Generating tokens for backend use
To generate tokens for headless clients, use  TokenMint.go:
```
go run src/authn/mint/TokenMint.go -key <a file containing a minting key> -user <a username the key should be assigned to>
```

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

#tests only need to be downloaded, not installed, need a special -t flag
go get -d -t fox/fox_test
go get -d -t login/login_test
```

## Running an Apache DS server

1. Install Apache DS and Apache Directory Studio
2. Import the config from FoxRegistryLDAPConfig.ldif and FoxRegistryLDAPConfig_system.ldif
3. Alternatively, as a hot fix, ldapProvider can be replaced with simpleProvider in authz.go

The travis file needs to be fixed for the build to succeed (LDAP server needs to put up on Travis' server).

Redhat's 389 DS should be used instead of Apache DS, since the latter's functionality is not that well documented.

## Running a REST server

1. Create folder $USER (current system username) in `config`, then copy and adapt example configuration file. 
2. Execute Fox binary passing an instance name as a parameter.

<<<<<<< HEAD

```
mkdir -p bin/config/$USER
cp src/config/config.json.template bin/config/$USER/config.json
cp src/config/config.json.template bin/config/$USER/test_config.json

mkdir /tmp/foxdb  # make sure that the configured storage folder exists.
=======

```
mkdir -p config/$USER
cp config/config.json.template config/$USER/config.json   # Normal config file
cp config/config.json.template config/$USER/test_config.json   # Config file for tests

mkdir /tmp/foxdb   # make sure that the configured storage folder exists.
>>>>>>> 21b12734e3a486bde2e91db6bae749b3ed1b0453
./bin/foxservice

go run src/authn/keygen/KeyGen.go > config/$USER/key.base64   # Generate the keyfile for authentication tokens
./bin/loginservice
```

REST interface will respond on **http://localhost:8090/**. You should now be able to use web UI in **http://localhost:9000/**.
To change a port used or logging target (defaults to stdout and can be sent to syslog), check **./bin/foxservice -h** and **./bin/loginservice -h**.

## Configuration
Configuration is user-based, every user has a folder with their username under `config/`, where their personal config file(s) live.
All services and tests use the same configuration file: `config.{ext}` for services, `test_config.{ext}` for tests. Config files can be in all formats supported by [Viper](http://github.com/spf13/viper) (JSON, TOML, YAML, HCL, Java properties).

All files referenced in configuration (keys, password list etc) should also be placed in user's config folder.
However, it's also possible to have both config and key files in config root folder as a fallback. These files are only used if corresponding file is not found in user's own folder. 

Order of loading config and referenced files:

1. User's folder `config/$USER/`
2. General config folder `config/` 

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
go run src/authn/pwd/pwdmaker/pwdMaker.go -user <username> -pwd <password> >> config/$USER/pwd.list
```

The pwd.list is a file referred to by the authn.PwdProvider.PwdFileName key in the config

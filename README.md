# Fox [![Build Status](https://travis-ci.org/e-gov/fox.svg?branch=master)](https://travis-ci.org/e-gov/fox)
An implementation of Fox Registry, the reference architecture for cloud-ready government systems. More detail of the project, its architecture and rationale is available at https://www.ria.ee/riigiarhitektuur/wiki/doku.php?id=an:rebasteregister

The FoxAPI application implements [this specification](http://editor.swagger.io/#/?import=https:%2F%2Fraw.githubusercontent.com%2Fe-gov%2Ffox%2Fmaster%2Ftatic%2F_data%2FFoxAPI.yaml)

There are two key components:
 1. The FoxService that implements the main business logic of the fox registry
 2. The LoginService that mints tokens for FoxService instances to be used and communicates with external authentication providers

## The microservices and their ports

Fox service's REST interface will respond on **http://localhost:8090/**, Login service's on **http://localhost:8091/**. You should now be able to use web UI in **http://localhost:9000/**. To change a port used or logging target (defaults to stdout and can be sent to syslog), check **./bin/foxservice -h** and **./bin/loginservice -h**.

**In the sample "config/pwd.list" file, a user FantasticMrFox is present with the encrypted version of "test".**

## Running the application

There are two ways to run this application:
1) Using Docker containers
2) Using bash

Scroll down to see the instructions on how to do this.

### The Docker way

It is highly recommended to use a Linux distrubution to run the Docker containers.

1) Install Docker
2) To avoid having to use the `sudo` command in front of the `docker` command, [create a docker user group](https://docs.docker.com/engine/installation/linux/ubuntulinux/#/create-a-docker-group)
3) Install Docker Compose
4) Run the app with `docker-compose up`

### The bash way
You can either run the build.sh file inside the project root, or do this manually.

#### 1 Starting a web-based UI

A web UI is built with `grunt`, to start it:

1. Make sure node.js and npm is installed.
2. Go to static folder `cd static`.
3. Run following commands

```bash
npm install
npm install -g grunt-cli
```
This should install all required tools to start the UI
Now go and copy `static/properties.json.sample` to `static/properties.json`
If you have default settings then sample properties will do just fine. Otherwise dig into properties.json

Now you should be all set. 
To run web ui run following command in `/static`
```bash
grunt serve
```
It should run webserver in `localhost:9000`

##### If you see errors about the encoding of files on OS X, try this:

```bash
export LC_ALL=en_US.UTF-8
export LANG=en_US.UTF-8
```

#### 2 Building a demo REST server

1. Change to the directory where the repository is cloned.
2. Setup environment and build application:

```bash
export GOPATH=$PWD
go get fox/foxservice
go get login/loginservice

# test dependencies only need to be downloaded,
# the test itself does not need to be installed.
# It also needs a special -t flag
go get -d -t fox/fox_test
go get -d -t login/login_test
```

#### 3 (optional) Running a LDAP server on Apache DS

A Director Service is, by default, not in use (see config.json.template in the config folder). However, one can be used, like this:
1. Install Apache DS and Apache Directory Studio
2. Create a new server in Apache DS (use the default configuration)
3. Run the server
4. Create a connection (again, use the default configuration)
5. Connect with the server
6. Import the config from FoxRegistryLDAPConfig_directories.ldif and FoxRegistryLDAPConfig_foxapi_as_user.ldif into the server
7. Change the Provider in config.json to "ldap"

Redhat's 389 DS should be used to replace Apache DS, since the latter's functionality is not that well documented.

#### 4 Running a REST server

1. Create folder $USER (current system username) in `config`, then copy and adapt example configuration file. 
2. Execute Fox binary passing an instance name as a parameter.

```bash
mkdir -p config/$USER
cp config/config.json.template config/$USER/config.json   # Normal config file
cp config/config.json.template config/$USER/test_config.json   # Config file for tests

mkdir /tmp/foxdb   # make sure that the configured storage folder exists.
./bin/foxservice

go run src/authn/keygen/KeyGen.go > config/$USER/key.base64   # Generate the keyfile for authentication tokens
./bin/loginservice
```

## Configuration
Configuration is user-based, every user has a folder with their username under `config/`, where their personal config file(s) live.
All services and tests use the same configuration file: `config.{ext}` for services, `test_config.{ext}` for tests. Config files can be in all formats supported by [Viper](http://github.com/spf13/viper) (JSON, TOML, YAML, HCL, Java properties).

All files referenced in configuration (keys, password list etc) should also be placed in user's config folder.
However, it's also possible to have both config and key files in config root folder as a fallback. These files are only used if corresponding file is not found in user's own folder. 

Order of loading config and referenced files:

1. User's folder `config/$USER/`
2. General, fallback, config folder `config/` 

## Reloading configuration
To reload configuration, both the login and fox services accept a HUP signal that should have both produce log messages about re-loading configuration

## Generating tokens for backend use
To generate tokens for headless clients, use  TokenMint.go:

```bash
go run src/authn/mint/TokenMint.go -key <a file containing a minting key> -user <a username the token should be assigned to>
```

## Generating passwords for authentication
To use the basic password authentication provider, passwords must be hashed and strored on server side. This happens like so:

```
touch pwd.list
go run src/authn/pwd/pwdmaker/pwdMaker.go -user <username> -pwd <password> >> config/$USER/pwd.list
```

The pwd.list is a file referred to by the authn.PwdProvider.PwdFileName key in the config.

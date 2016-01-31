# Fox
An implementation of Fox Registry, the reference architecture for cloud-ready government systems. More detail of the project, its architecture and rationale is available at https://www.ria.ee/riigiarhitektuur/wiki/doku.php?id=an:rebasteregister

The FoxAPI application implements [this specification](http://editor.swagger.io/#/?import=https:%2F%2Fraw.githubusercontent.com%2Fe-gov%2Ffox%2Fmaster%2Ftatic%2F_data%2FFoxAPI.yaml)

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
go get fox      # get dependencies
go install fox  # application will be built into bin/fox
```

## Running a REST server

1. Copy and adapt example configuration file:
2. Execute Fox binary passing an instance name as a parameter.

```
cp src/fox/config.gcfg.template bin/config.gcfg
mkdir /tmp/foxdb  # make sure that the configured storage folder exists.
cd bin
./fox my
```


REST interface will respond on **http://localhost:8090/**. You should now be able to use web UI.

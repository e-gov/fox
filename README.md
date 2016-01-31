# Fox
An implementation of Fox Registry, the reference architecture for cloud-ready government systems. More detail of the project, its architecture and rationale is available at https://www.ria.ee/riigiarhitektuur/wiki/doku.php?id=an:rebasteregister

The FoxAPI application implements [this specification](http://editor.swagger.io/#/?import=https:%2F%2Fraw.githubusercontent.com%2Fe-gov%2Ffox%2Fmaster%2Ftatic%2F_data%2FFoxAPI.yaml)

The static site is built for jekyll, use `jekyll serve` to launch

## Building a demo application

1. Change to the directory where the repository is cloned.
2. Setup environment and build application:

```
export GOPATH=$PWD
go get fox  # get dependencies
go install fox # application will be built into bin/fox
```

## Running an application

1. Copy and adapt example configuration file:

```
cp src/fox/config.gcfg.template bin/config.gcfg
mkdir /tmp/foxdb  # make sure that the configured storage folder exists.
```

2. Execute Fox binary (**bin/fox**) passing an instance name as a parameter.

REST interface will respond on **http://localhost:8090/**.

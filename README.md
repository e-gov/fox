# Fox
An implementation of Fox Registry, the reference architecture for cloud-ready government systems. More detail of the project, its architecture and rationale is available at https://www.ria.ee/riigiarhitektuur/wiki/doku.php?id=an:rebasteregister

The FoxAPI application implements [this specification](http://editor.swagger.io/#/?import=https:%2F%2Fraw.githubusercontent.com%2Fe-gov%2Ffox%2Fmaster%2Ftatic%2F_data%2FFoxAPI.yaml)

The static site is built for jekyll, use `jekyll serve` to launch

To compile the GO app do
```
go install
go build
```

After which you end up with a binary that takes a single parameter (the instance name) to launch

#!/bin/bash

cd static
npm install
npm install -g grunt-cli
cp properties.json.sample properties.json
grunt serve

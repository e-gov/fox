#!/bin/bash

#https://lebkowski.name/docker-volumes/

docker ps --filter status=dead --filter status=exited -aq | xargs -r docker rm -v

docker images --no-trunc | grep '<none>' | awk '{ print $3 }' | xargs -r docker rmi

docker images --no-trunc | grep '<none>' | awk '{ print $3 }' | xargs -r docker rmi

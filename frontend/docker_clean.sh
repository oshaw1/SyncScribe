#!/bin/bash

# Delete all containers

docker container stop $(docker container ls -aq)
docker container rm $(docker container ls -aq)
docker system prune -a -f --volumes
#!/bin/bash

clear_docker_containers() {
  for F in `docker ps -a -q -f status=exited`; do
    docker rm -v $F
  done
}

clear_docker_images() {
  for F in `docker images -f "dangling=true" -q`; do
    docker rmi $F
  done
}

#!/bin/bash

clear_docker_containers() {
  docker ps -aqf status=exited | xargs -I {} docker rm {}
}

clear_docker_images() {
  docker images -qf "dangling=true" -q | xargs -I {} docker rmi {}
}

docker_purge_container() {
  NAME=${1?docker name}
  docker ps \
    -a \
    --format '{{.ID}}\t{{.Names}}' \
    | awk \
    -v name=$NAME \
    '$2==name{system("docker stop "$1" && docker rm "$1)}'
}


#!/bin/bash

export DOCKER_BUILDKIT=1

function clear_docker_containers() {
  docker ps -aqf status=exited | xargs -I {} docker rm {}
}

function clear_docker_images() {
  docker images -qf "dangling=true" -q | xargs -I {} docker rmi {}
}

function docker_purge_container() {
  NAME=${1?docker name}
  docker ps \
    -a \
    --format '{{.ID}}\t{{.Names}}' \
    | awk \
    -v name=$NAME \
    '$2==name{system("docker stop "$1" && docker rm "$1)}'
}

function docker_mysql() {

  # When initiating use the following parameter termplate to specify root password.
  #  -e MYSQL_ROOT_PASSWORD="<generated root password>" \
  docker run \
    -d --restart unless-stopped \
    --name "mysql8" \
    -p 33060:3306 \
    -v $HOME/data/container/mysql8/var/lib/mysql:/var/lib/mysql \
    mysql:8.0 \
    --default-authentication-plugin=mysql_native_password \

}

function docker_clean_build_cache() {
  docker builder prune
}

# docker_mysql

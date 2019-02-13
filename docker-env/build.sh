#!/bin/bash

set -ex

DOCKERFILE_DIR=${1?Dockerfile dir}
if [ -d "$DOCKERFILE_DIR" ]; then
  BASE_DIR=$DOCKERFILE_DIR
else
  if [ "$(basename $DOCKERFILE_DIR)" == "Dockerfile" ]; then
    BASE_DIR=$(dirname "$DOCKERFILE_DIR")
  else
    echo "Must specify the directory containing Dockerfile"
    exit 1
  fi
fi

bindir=$(dirname $0)
bindir=$(cd $bindir;pwd)

BASE_DIR=$(cd $BASE_DIR;pwd)
APP_NAME=$(basename $BASE_DIR)

docker build \
  -t ${APP_NAME}:latest \
  $BASE_DIR

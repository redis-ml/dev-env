#!/bin/bash

set -ex

bindir=$(dirname $0)
bindir=$(cd $bindir;pwd)
APP_NAME=$(basename $bindir)

docker build \
  -t ${APP_NAME}:latest \
  .

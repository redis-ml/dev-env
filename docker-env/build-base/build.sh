#!/bin/bash

bindir=$(dirname $0)
bindir=$(cd $bindir;pwd)

cd $bindir

docker build -t docker-base:latest .

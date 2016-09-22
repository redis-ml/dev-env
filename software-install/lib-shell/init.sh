#!/bin/bash

bindir=`dirname $0`
bindir=`cd $bindir;pwd`

mkdir ~/.env/ \
  && virtualenv ~/.env/default \
  && rsync -a $bindir/lib-shell ~/.env/ \
  && rsync $bindir/default.sh ~/.env/

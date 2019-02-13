#!/bin/bash

### Common logic for all scripts.
set -ex

bindir=$(dirname $0)
bindir=$(cd $bindir;pwd)
TOP_DIR=$(cd $bindir/../../;pwd)
echo top dir: $TOP_DIR
### [ENV] Common logic for all scripts.

INSTALL_LIB=$TOP_DIR/software-install/libs
#########
. $INSTALL_LIB/_env.sh $INSTALL_LIB

install_pyenv
install_rbenv
install_golang


#!/bin/bash

export MY_ENV_SCRIPT_DIR=~/.env/default
export MY_SOFTWARE_DIR=~/.local

# add local bin dir
export PATH=$PATH:$MY_SOFTWARE_DIR/bin

for F in `find ${MY_ENV_SCRIPT_DIR}/ -maxdepth 1 -type f -name "lib.sh" -prune -o  -type f -name "*.sh" -print`; do
    echo $F
    . $F
done

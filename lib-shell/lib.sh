#!/bin/bash

MY_LIB_PATH=~/lib-shell

for F in `find ${MY_LIB_PATH}/ -name "lib.sh" -prune -o -type f -name "*.sh" -print`; do
    echo $F
    . $F
done


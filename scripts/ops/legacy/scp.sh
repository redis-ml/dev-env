#!/bin/bash

. func-ssh.sh


if [ $# -lt 1 ]; then
    usage
    exit 1
fi

MY_SOURCE=$1
shift
MY_DEST=$1
shift

parse_arg "$@"

# replace customized keywords
MY_SOURCE=${MY_SOURCE/#remote:/ubuntu@${MY_HOST}:}
MY_DEST=${MY_DEST/#remote:/ubuntu@${MY_HOST}:}

MY_CMD="scp -r -i ${MY_PEM_FILE}  $MY_SOURCE $MY_DEST"

echo ${MY_CMD}
sh -c "${MY_CMD}"


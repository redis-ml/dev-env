#!/bin/bash

. func-ssh.sh

if [ $# -lt 1 ]; then
    usage
    exit 1
fi

EXTRA_ARGS=""
case "$1" in
    screen)
        EXTRA_ARGS="screen -r lmz -d"
        shift
        ;;
esac

parse_arg "$@"

MY_CMD="ssh -l ubuntu ${MY_HOST} -i ${MY_PEM_FILE} -t ${EXTRA_ARGS}"

echo ${MY_CMD}
sh -c "${MY_CMD}"

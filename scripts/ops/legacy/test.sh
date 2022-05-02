#!/bin/bash

set -ex
    docker run \
      --rm -ti \
            -v /etc/ssl/certs:/etc/ssl/certs \
            -v $HOME/.ssh:/root/.ssh \
            -p 127.0.0.1:8888:8888 \
            ssh-client \
            ssh -N -D 0.0.0.0:8888 dev.redis.biz


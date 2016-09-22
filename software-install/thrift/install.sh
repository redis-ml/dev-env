#!/bin/bash

sudo apt-get install automake bison flex g++ git libboost1.55-all-dev libevent-dev libssl-dev libtool make pkg-config

./bootstrap.sh \
    && ./configure \
         --without-java \
    && make \
    && make check \
    && echo make install


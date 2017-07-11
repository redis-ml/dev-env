#!/bin/bash


CACHE_DIR=~/github-repo/downloads

save_with_local_cache() {
    local url=$1
    local FN=$2
    wget -N -c -O "$FN" "$url"
}

set -ex

install_base() {
  sudo apt-get install \
      g++ \
      automake \
      autoconf \
      autoconf-archive \
      libtool \
      libboost-all-dev \
      libevent-dev \
      libdouble-conversion-dev \
      libgoogle-glog-dev \
      libgflags-dev \
      liblz4-dev \
      liblzma-dev \
      libsnappy-dev \
      make \
      zlib1g-dev \
      binutils-dev \
      libjemalloc-dev \
      libssl-dev

  sudo apt-get install \
      libunwind8-dev \
      libelf-dev \
      libdwarf-dev

  sudo apt-get install \
      libiberty-dev
}

install_libgtest() {
  echo install_libgtest
  pwd
  [ -d googletest ] || git clone https://github.com/google/googletest.git
  #########################

  sudo apt-get install cmake # install cmake

  cd googletest

  git checkout release-1.7.0 || echo "falied"

  sudo cmake CMakeLists.txt

  sudo make

  sudo rsync *.a /usr/local/lib/
  sudo rsync -a include/gtest /usr/local/include/
  cd -
}
install_folly() {
  #########################
  echo install_folly
  pwd
  [ -d folly ] || git clone https://github.com/facebook/folly.git

  TOP_DIR=$PWD
  cd folly/folly

    cd test/
    local DST_DIR=gtest-1.7.0
    rm -rf "${DST_DIR}" "${DST_DIR}.zip"
    rsync -az $TOP_DIR/googletest $DST_DIR
    tar czf gtest-1.7.0.zip $DST_DIR
    cd ../

    autoreconf -ivf
    ./configure
    make clean
    make
    make check || echo "Small amounts of failure could be ignored, please check..."
    sudo make install
  cd ../..
}

install_proxygen_step1() {
  [ -d proxygen ] || git clone https://github.com/facebook/proxygen.git
  echo install_proxygen_step1
  #########################
  pwd
  cd proxygen/proxygen
  sh -c ./deps.sh || GOOD=1
  cd -
}

install_fbthrift_py() {
  #########################
  [ -d fbthrift ] || git clone https://github.com/facebook/fbthrift.git
  pwd
  cd fbthrift/thrift
  ./build/deps_ubuntu_14.04.sh
    cd lib/py
    sudo python setup.py install
    cd -
  cd ../..
}

install_wangle() {
  [ -d wangle ] || git clone https://github.com/facebook/wangle.git
  pwd
  #########################
  cd wangle/wangle
  cmake .
  cd -
  cd folly/folly
  sudo make install
  cd -
  cd wangle/wangle
  make
  sudo make install
  cd -
}

install_proxygen() {
  [ -d proxygen ] || git clone https://github.com/facebook/proxygen.git
  #########################
  pwd
  TOP_DIR="$PWD"
  cd proxygen/proxygen

  rm -rf folly
  ln -s ../../folly folly
  rm -rf Makefile
  rm -rf config.log
  rm -rf config.status

  # Copy gtest from folly
  # download from "https://codeload.github.com/google/googlemock/zip/release-1.7.0"
  perl -p -i -e 's#wget https://googlemock.googlecode.com/files/gmock-1.7.0.zip#wget https://codeload.github.com/google/googlemock/zip/release-1.7.0 -O gmock-1.7.0.zip' ./lib/test/Makefile.am
  #rsync $TOP_DIR/folly/folly/test/gtest-1.7.0.zip lib/test/

  autoreconf -ivf
  ./configure
  ./reinstall.sh

  cd ../../
}

install_fbthrift() {
  sudo apt-get install make autoconf libtool g++ \
    libboost-all-dev libevent-dev flex bison \
    libgoogle-glog-dev libdouble-conversion-dev scons \
    libkrb5-dev libsnappy-dev libsasl2-dev
  [ -d fbthrift ] || git clone https://github.com/facebook/fbthrift.git

  # specific version
  git checkout v2016.12.05.00

  #########################
  pwd
  cd fbthrift/thrift

  bash ./build/deps_ubuntu_14.04.sh
  autoreconf -if && \
    PY_PREFIX=/ ./configure --without-py && \
    make
  sudo make install

  cd -
}

install_rocksdb() {
  #########################
  sudo apt-get install libgflags-dev
  sudo apt-get install libsnappy-dev
  sudo apt-get install zlib1g-dev
  sudo apt-get install libbz2-dev
  [ -d rocksdb ] || git clone https://github.com/facebook/rocksdb.git
  pwd
  cd rocksdb
  make shared_lib
  sudo make install
  cd -
}

install_others() {
  #########################
  sudo apt-get install libcurl4-openssl-dev

  sudo apt-get install libhiredis-dev
}

main() {
  install_base
  install_libgtest
  install_folly
  install_proxygen_step1
  install_fbthrift_py
  install_wangle
  // Install folly again.
  install_proxygen
  install_fbthrift
  install_rocksdb
  install_others
  ########################
  echo "Download latest zookeeper from:"
  echo "http://www-us.apache.org/dist/zookeeper/stable/"
  echo
  echo "then run the following cmds:"
  
  echo "cd zookeeper-3.4.8/src/c"
  echo "./configure && make"
  echo "sudo make install"
  echo ""
  echo "cp conf/zoo_sample.cfg conf/zoo.cfg"
  echo "bin/zkServer.sh start"
}


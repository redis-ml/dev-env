#!/bin/bash

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

install_folly_step1() {
  #########################
  echo install_folly_step1
  pwd
  [ -d folly ] || git clone https://github.com/facebook/folly.git

  cd folly/folly

    cd test/gtest-1.7.0
    cmake CMakeLists.txt
    make
    cd -

    autoreconf -ivf
    ./configure
    make clean
    make
    make check || echo "Small amounts of failure could be ignored, please check..."
    sudo make install
  cd ../..
}

install_libgtest() {
  echo install_libgtest
  pwd
  [ -d googletest ] || git clone https://github.com/google/googletest.git
  #########################

  sudo apt-get install cmake # install cmake

  cd googletest

  sudo cmake CMakeLists.txt

  sudo make

  sudo cp *.a /usr/local/lib/
  sudo cp -r include/gtest /usr/local/include/
  cd -
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

install_fbthrift() {
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

install_folly() {
  [ -d folly ] || git clone https://github.com/facebook/folly.git
  #########################
  pwd
  cd folly/folly
    make clean
    autoreconf -ivf
    ./configure
    make clean
    make
    make check || echo "it failed but doesn't matter"
    sudo make install
  cd -
}

install_proxygen() {
  [ -d proxygen ] || git clone https://github.com/facebook/proxygen.git
  #########################
  pwd
  cd proxygen/proxygen

  rm -rf folly
  ln -s ../../folly folly
  rm -rf Makefile
  rm -rf config.log
  rm -rf config.status
  autoreconf -ivf
  ./configure
  ./reinstall.sh

  cd -
}

install_fbthrift_thrift() {
  [ -d fbthrift ] || git clone https://github.com/facebook/fbthrift.git
  #########################
  pwd
  cd fbthrift/thrift

  autoreconf -if && ./configure && make
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
  install_folly_step1
  install_proxygen_step1
  install_fbthrift
  install_wangle
  install_folly
  install_proxygen
  install_fbthrift_thrift
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


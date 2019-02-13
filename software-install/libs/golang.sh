#!/bin/bash

install_golang() {
  local install_path=${1:-/opt}
  # Golang
  local GO_VERSION="1.11.4"
  local INSTALL_DIR="go${GO_VERSION}.linux-amd64"
  local GOLANG_DFILE="${INSTALL_DIR}.tar.gz"
  [ -f ${install_path}/$GOLANG_DFILE ] || curl https://storage.googleapis.com/golang/$GOLANG_DFILE > ${install_path}/$GOLANG_DFILE
  [ -d ${install_path}/${INSTALL_DIR} ] || mkdir ${install_path}/${INSTALL_DIR}
  tar xf ${install_path}/$GOLANG_DFILE -C ${install_path}/${INSTALL_DIR} --strip-components 1
  ln -sf ${install_path}/${INSTALL_DIR} ${install_path}/go
}

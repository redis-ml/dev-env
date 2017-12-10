#!/bin/bash

export GOROOT=${MY_SOFTWARE_DIR}/go

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

go_install() {
  local SUB=${1?Need package relative path}
  go install ${MY_PACKAGE_NAME}/$SUB
}

go_test() {
  local SUB=${1?Need package relative path}
  go test ${MY_PACKAGE_NAME}/$SUB
}

go_run() {
  local SUB=${1?Need package relative path}
  go run ${MY_PACKAGE_NAME}/$SUB
}


#!/bin/bash

export GO111MODULE=on

go_install() {
  local SUB=${1?Need package relative path}
  go install -ldflags="-X main.Version=$(git describe --always --long --dirty)" ${MY_PACKAGE_NAME}/$SUB
}

go_test() {
  local SUB=${1?Need package relative path}
  go test ${MY_PACKAGE_NAME}/$SUB
}

go_run() {
  local SUB=${1?Need package relative path}
  go run ${MY_PACKAGE_NAME}/$SUB
}


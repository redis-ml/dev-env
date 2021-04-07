#!/bin/bash

export GO111MODULE=on

function go_install() {
  local SUB=${1?Need package relative path}
  go install -ldflags="-X main.Version=$(git describe --always --long --dirty)" ${MY_PACKAGE_NAME}/$SUB
}

function go_test() {
  local SUB=${1?Need package relative path}
  go test ${MY_PACKAGE_NAME}/$SUB
}

function go_run() {
  local SUB=${1?Need package relative path}
  go run ${MY_PACKAGE_NAME}/$SUB
}

GO_BIN_PATH="$(go env GOPATH)/bin"
( echo $PATH | fgrep "$GO_BIN_PATH:" 2>&1 >/dev/null ) || {
  which go >/dev/null 2>/dev/null && export PATH="$(go env GOPATH)/bin:$PATH"
}

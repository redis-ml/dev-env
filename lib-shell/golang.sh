#!/bin/bash


GOROOT="$(go env GOROOT 2>/dev/null)"

if [ "$GOROOT" = "" ]; then
  export GOROOT=/usr/local/opt/go/libexec
fi

GOPATH="$(go env GOPATH 2>/dev/null)"
if [ "$GOPATH" = "" ]; then
  GOPATH="$(cd ~/go;pwd)"
fi

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

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


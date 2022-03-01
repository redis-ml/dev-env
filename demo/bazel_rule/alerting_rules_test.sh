#!/bin/bash

echo pwd: $PWD
echo $TEST_SRCDIR
ls -l
cd "$TEST_SRCDIR"
echo pwd: $PWD

PROMTOOL_BIN=$(ls prometheus_prometool_*/prometheus*/promtool)
$PROMTOOL_BIN --help


# ./prometheus_prometool_macos/prometheus-2.27.1.darwin-amd64/promtool
exit 1

#!/bin/bash

THIS_DIR=${1?real dir of this file}

. $THIS_DIR/_common.sh

for f in `ls $THIS_DIR/*.sh`; do
  echo "checking $f"
  BASE_NAME="$(basename $f)"
  if [[ "$BASE_NAME" =~ ^_ ]]; then
    echo "skiping $f"
    continue
  fi

  . $f
done


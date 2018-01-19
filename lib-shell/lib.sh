#!/bin/bash

# Initialization
export MY_OLD_PATH=${MY_OLD_PATH:-$PATH}
export PATH=${MY_OLD_PATH}

if [ "x${MY_ENV_SCRIPT_DIR}" != "x" ]; then
  for F in `find ${MY_ENV_SCRIPT_DIR}/ -maxdepth 1 -type f -name "lib.sh" -prune -o  -type f -name "*.sh" -print`; do
      echo $F
      . $F
  done
else
  echo "!!!  !!!   !!!   !!!"
  echo "you MUST evaluate default.sh first"
  echo "exiting without anything"
fi


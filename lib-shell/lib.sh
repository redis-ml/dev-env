#!/bin/bash

# Initialization
export MY_OLD_PATH=${MY_OLD_PATH:-$PATH}
export PATH=${MY_OLD_PATH}

if [ "${MY_ENV_SCRIPT_DIR}" != "" ]; then
  for F in `find ${MY_ENV_SCRIPT_DIR} -maxdepth 1 -type f -name "lib.sh" -prune -o  -type f -name "*.sh" -print`; do
      if [ "MY_DEV_ENV_DEBUG" = "true" ]; then
          echo $F
      fi
      . $F
  done
else
  echo "!!!  !!!   !!!   !!!"
  echo "you MUST evaluate default.sh first"
  echo "exiting without anything"
fi


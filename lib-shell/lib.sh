#!/bin/bash

# Initialization
export MY_OLD_PATH=${MY_OLD_PATH:-$PATH}
export PATH=${MY_OLD_PATH}

if [ "${MY_ENV_SCRIPT_DIR}" != "" ]; then
  for F in `find ${MY_ENV_SCRIPT_DIR} -maxdepth 1 -type f -name "lib.sh" -prune -o -type f -name "*.sh" -print`; do
      if [ "$MY_DEV_ENV_DEBUG" = "true" ]; then
          echo $F
      fi
      source $F
  done
else
  echo "!!!  !!!   !!!   !!!"
  echo "you MUST evaluate default.sh first"
  echo "exiting without anything"
fi

# Make sure MacPorts is in local path
export PATH="$HOME/local/bin:$PATH"

# Tool func to quickly get into this dicrectory
alias cd_dev_env_tools="cd $MY_ENV_SCRIPT_DIR"


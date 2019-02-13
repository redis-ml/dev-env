#!/bin/bash

install_pyenv() {
  local install_path=${1:-/opt}
  export PYENV_ROOT="${install_path}/pyenv"
  clone_repo https://github.com/pyenv/pyenv.git $PYENV_ROOT
  export PATH="$PYENV_ROOT/bin:$PATH"
  eval "$(pyenv init -)"
  pyenv install 2.7.14
  pyenv install 3.6.8
}

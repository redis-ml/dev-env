#!/bin/bash

# Ruby/rbenv
install_rbenv() {
  local install_path=${1:-/opt}
  # This is must
  export RBENV_ROOT="${install_path}/rbenv"
  clone_repo https://github.com/rbenv/rbenv.git $RBENV_ROOT
  export PATH="$RBENV_ROOT/bin:$PATH"
  local RBENV_PLUGIN_DIR="$(rbenv root)"/plugins
  [ -d $RBENV_PLUGIN_DIR ] || mkdir -p $RBENV_PLUGIN_DIR
  clone_repo https://github.com/rbenv/ruby-build.git "$(rbenv root)"/plugins/ruby-build
  eval "$(rbenv init -)"
  rbenv install 2.3.1
  rbenv install 2.6.1
}


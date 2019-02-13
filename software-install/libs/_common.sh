#!/bin/bash

clone_repo() {
  local url=${1?repo url}
  local dir=${2?local dir}

  [ -d "$dir" ] || git clone "$url" "$dir"
}


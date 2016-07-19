#!/bin/bash

git_help() {
    echo -e "Use the following command for git operations:\n\n"
    echo -e "   git_push_to_remote : push local branch changes to remote\n\n"
    echo -e "   git_commit : commit changes to previous commit\n\n"
}

git_push_to_remote() {
  git push origin HEAD:master
}

git_commit() {
  git commit --amend "$@"
}

git_repo_clone() {
    local url=${1?}
    local release=${2:-unset}

    local P="${url/*\//}"
    local dir=${P/%.git/}
    [ -d $dir ] && return
    git clone $url
    pushd $dir
    if [ "$release" != "unset" ]; then
        git checkout -b $release $release 
    fi
    popd
}

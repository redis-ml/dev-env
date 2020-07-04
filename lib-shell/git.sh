#!/bin/bash

git_help() {
    echo -e "Use the following command for git operations:\n\n"
    echo -e "   git_push_to_remote : push local branch changes to remote\n\n"
    echo -e "   git_commit : commit changes to previous commit\n\n"
}

git_fetch() {
  git fetch
}
git_rebase() {
  git rebase origin/master
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

# This is very slow in Monorepo.
#export GIT_PS1_SHOWDIRTYSTATE=1

# this relies on git-prompt.sh
# 
[ -f $MY_ENV_SCRIPT_DIR/git-prompt.sh ] \
  || curl -L \
    https://raw.github.com/git/git/master/contrib/completion/git-prompt.sh \
    > $MY_ENV_SCRIPT_DIR/git-prompt.sh 
# I'm using 'bash_it_theme/mingzhu.theme.bash' instead.
# export PS1='\w$(__git_ps1 " (%s)")\$ '

export EDITOR=vim

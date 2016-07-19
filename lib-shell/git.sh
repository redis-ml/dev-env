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


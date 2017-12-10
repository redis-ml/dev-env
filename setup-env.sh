#!/bin/bash

bindir=`dirname $0`
bindir=`cd $bindir;pwd`

git config --global user.name "Mingzhu Liu"
git config --global user.email "redisliu@gmail.com"

git config --global core.autocrlf input
git config --global core.editor vim

# Seamlessly integrate Go tool and private repos.
git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
# Eliminate Git 2.0 warnings.
git config --global push.default simple

# Install Golang plugin.
sh $bindir/software-install/vim/init.sh

# Generate ssh-key if needed.
[ -f ~/.ssh/id_rsa.pub ] || ssh-keygen

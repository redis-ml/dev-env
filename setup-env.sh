#!/bin/bash

bindir=`dirname $0`
bindir=`cd $bindir;pwd`

git config --global user.name "Mingzhu Liu"
git config --global user.email "redisliu@gmail.com"

git config --global core.autocrlf input
git config --global core.editor vim

sh $bindir/software-install/vim/init.sh

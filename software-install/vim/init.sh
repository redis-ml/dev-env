#!/bin/bash

bindir=`dirname $0`
bindir=`cd $bindir;pwd`

mkdir -p ~/.vim/autoload ~/.vim/bundle \
 && curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim \
 && cp $bindir/vimrc ~/.vimrc \
 && git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go


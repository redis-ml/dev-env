#!/bin/bash

bindir=`dirname $0`
bindir=`cd $bindir;pwd`

# mkdir -p ~/.vim/autoload ~/.vim/bundle \
#  && curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim \
echo "install vim" \
 && cp $bindir/vimrc ~/.vimrc \
 && ( \
   [ -d ~/.vim/pack/plugins/start/vim-go ] \
   || git clone https://github.com/fatih/vim-go.git \
      ~/.vim/pack/plugins/start/vim-go \
    ) \
 && ( \
   [ -d ~/.vim/pack/dist/start/nerdtree ] \
    || git clone https://github.com/scrooloose/nerdtree.git \
      ~/.vim/pack/dist/start/nerdtree \
    ) \


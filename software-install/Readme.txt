Each of every subfolder contains several scripts for installing/initiating corresponding software. For now, there are 2 kinds of scripts:

* install.sh

  Such script installs software into system level directories, like php/python/etc, which needs to be install in something like /usr/bin, /usr/local/bin/, /usr/lib.

* init.sh

  Such scripts fetch the software into ***Current Directory***, instead of installing it into some common places. Such softwares are usually on case-by-case basis.
  For example, 'Composer' initiates the directory for a new PHP project.

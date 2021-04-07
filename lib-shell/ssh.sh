#!/bin/bash
export ITERM_ENABLE_SHELL_INTEGRATION_WITH_TMUX=YES

function my_usage() {
    echo "Usage:"
    echo "      sh ~/ssh.sh <host ip or config name> [category]"
    echo "      sh ~/scp.sh <src_dir> <dest_dir> <host ip or config name> [category]"
    echo
    echo
    echo "  known args for 'ssh.sh':"
    echo "    sh ~/ssh.sh default"
    echo "    sh ~/ssh.sh 152.4.17.7 cate"
    echo "  could also have 'screen' before those parameters:"
    echo "    sh ~/ssh.sh screen default"
    echo "  which will launch 'screen -r lmz -d' upon connected"
    echo
    echo "  known args for scp.sh:"
    echo "    sh ~/scp.sh ./abc remote:~/something default"
    echo "    sh ~/scp.sh ./abc remote:~/something default"
    echo "    sh ~/scp remote:~/abc /tmp/ 152.4.17.7 cate"
    echo "  'remote' is a keyword of this script which will be replaced later."
}

function parse_arg() {
  case $1 in
      default)
          MY_HOST=127.0.0.1
          MY_CATEGORY=cate
          ;;
      *)
          MY_HOST=${1}
          MY_CATEGORY=${2}
          ;;
  esac

  MY_PEM_DIR="~/.ssh/keys"

  case ${MY_CATEGORY} in
    cate)
      MY_PEM_FILE="${MY_PEM_DIR}/cate.pem"
      ;;
    *)
      usage
      exit 1
      ;;
  esac
}

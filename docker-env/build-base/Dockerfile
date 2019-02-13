# Env for building dev tools, including pyenv, rbenv, etc.

FROM ubuntu:latest

RUN \
  DEBIAN_FRONTEND='noninteractive' apt-get update && \
  DEBIAN_FRONTEND='noninteractive' apt-get install -y \
    vim groff curl tmux \
    apt-transport-https ca-certificates software-properties-common \
    build-essential git \
    libreadline-dev zlib1g-dev libssl-dev libbz2-dev libsqlite3-dev

# overwrite this with 'CMD []' in a dependent Dockerfile
CMD ["/bin/bash"]

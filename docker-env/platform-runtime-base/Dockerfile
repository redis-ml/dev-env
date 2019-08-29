# Base image for building anything.

FROM ubuntu:latest

RUN \
  DEBIAN_FRONTEND='noninteractive' apt-get update && \
  DEBIAN_FRONTEND='noninteractive' apt-get install -y \
    vim groff curl \
    apt-transport-https ca-certificates software-properties-common \
    git libreadline-dev zlib1g-dev libssl-dev libbz2-dev libsqlite3-dev \
    openssl \
    && apt-get clean && \
    echo "done"

ENV HOME /root
ENV PYENV_ROOT $HOME/.pyenv 
ENV RBENV_ROOT $HOME/.pyenv 
ENV PATH $PYENV_ROOT/bin:$RBENV_ROOT/bin:$PATH


# overwrite this with 'CMD []' in a dependent Dockerfile
CMD ["/bin/bash"]

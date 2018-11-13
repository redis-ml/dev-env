# AWS client enabled.

FROM ubuntu:latest

RUN \
  DEBIAN_FRONTEND='noninteractive' apt-get update && \
  DEBIAN_FRONTEND='noninteractive' apt-get install -y \
    vim groff curl \
    apt-transport-https ca-certificates software-properties-common \
    build-essential git libreadline-dev zlib1g-dev libssl-dev libbz2-dev libsqlite3-dev

# Add Docker repo
RUN \
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - && \
  add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# Install Docker
RUN \
  DEBIAN_FRONTEND='noninteractive' apt-get update && \
  DEBIAN_FRONTEND='noninteractive' apt-get install -y docker-ce

RUN curl -L https://raw.githubusercontent.com/pyenv/pyenv-installer/master/bin/pyenv-installer | bash

ENV LANG en_US.UTF-8
ENV PYENV_ROOT /.pyenv
ENV PATH $PYENV_ROOT/shims:$PYENV_ROOT/bin:$PATH

RUN pyenv install 3.6.6
RUN pyenv install 2.7.15

# overwrite this with 'CMD []' in a dependent Dockerfile
CMD ["/bin/bash"]

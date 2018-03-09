FROM ubuntu:latest

RUN \
  apt-get -y update 

RUN \
  apt-get -y install \
    python-pip \
  && \
  pip install awscli && \
  pip install virtualenv && \
  echo done

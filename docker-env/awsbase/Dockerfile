# AWS client enabled.

FROM ubuntu:latest

RUN \
  apt-get update && \
  apt-get install -y \
    python-pip vim groff curl \
    apt-transport-https ca-certificates software-properties-common
RUN \
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - && \
  add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" && \
  apt-get update && \
  apt-get install -y docker-ce

# overwrite this with 'CMD []' in a dependent Dockerfile
CMD ["/bin/bash"]

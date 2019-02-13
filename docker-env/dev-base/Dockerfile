# Dev env with popular language installed.

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

# overwrite this with 'CMD []' in a dependent Dockerfile
CMD ["/bin/bash"]

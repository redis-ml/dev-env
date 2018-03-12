# AWS client enabled.

FROM awsbase:latest

RUN apt-get install -y python-pip vim
RUN pip install awscli

# Deprecated.
# Docker not work inside containder.
# RUN apt-get install -y apt-transport-https ca-certificates curl software-properties-common awscli vim
# RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
# RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
# RUN apt-get update
# RUN apt-get install -y docker-ce

# overwrite this with 'CMD []' in a dependent Dockerfile
CMD ["/bin/bash"]

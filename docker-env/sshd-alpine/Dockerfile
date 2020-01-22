FROM alpine:latest

RUN apk add --no-cache openssh

RUN \
  mkdir -p /root/.ssh && \
  chown root:root -R /root/ && \
  chmod 0700 /root/ && \
  chmod 0700 /root/.ssh && \
  echo "done"


#RUN sed -i s/#PubkeyAuthentication.*/PubkeyAuthentication\ yes/ /etc/ssh/sshd_config && \
#    sed -i s/#PermitRootLogin.*/PermitRootLogin\ no/ /etc/ssh/sshd_config && \
#    sed -i s/PermitRootLogin.*/PermitRootLogin\ no/ /etc/ssh/sshd_config && \
#    sed -i s/#LogLevel.*/LogLevel\ DEBUG/ /etc/ssh/sshd_config && \
RUN \
  ssh-keygen -t rsa -b 4096 -f /etc/ssh/ssh_host_rsa_key  && \
  ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key  && \
  ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key  && \
  echo "done"
RUN \
  sed -i s/#GatewayPorts.*/GatewayPorts\ yes/ /etc/ssh/sshd_config && \
  echo "done"

RUN rm -rf /var/cache/apk/*

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]

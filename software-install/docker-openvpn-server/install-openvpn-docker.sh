#!/bin/bash

SERVER_DOMAIN=${1:-f.redis.biz}

OVPN_DATA="ovpn-data"

docker run --name $OVPN_DATA -v /etc/openvpn busybox

docker run --volumes-from $OVPN_DATA --rm kylemanna/openvpn ovpn_genconfig -u udp://${SERVER_DOMAIN}:1194

docker run --volumes-from $OVPN_DATA --rm -it kylemanna/openvpn ovpn_initpki

cat > /etc/init/docker-openvpn.conf <<EOF
description "Docker container for OpenVPN server"
start on filesystem and started docker
stop on runlevel [!2345]
respawn
script
  exec docker run --volumes-from ovpn-data --rm -p 1194:1194/udp --cap-add=NET_ADMIN kylemanna/openvpn
end script
EOF

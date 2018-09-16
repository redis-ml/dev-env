#!/bin/bash

set -ex

REPO=${APP_REPO:-usdcny}
APP_NAME=${APP_NAME:-openvpn-docker}

CID=$(docker ps -qf "name=^/${APP_NAME}\$")
if [ "$CID" != "" ]; then
  docker stop $CID
  docker rm $CID
fi

exec docker run \
  -d \
  --restart unless-stopped \
  --log-opt mode=non-blocking \
  --log-opt max-buffer-size=64m \
  --volumes-from ovpn-data \
  -p 1194:1194/udp \
  --cap-add=NET_ADMIN \
  --name $APP_NAME \
  kylemanna/openvpn

#!/bin/bash

set -ex

CLIENTNAME=${1:-MyMac}

OVPN_DATA="ovpn-data"

docker run \
  --volumes-from $OVPN_DATA \
  --rm -it \
  kylemanna/openvpn \
  easyrsa build-client-full $CLIENTNAME nopass

docker run \
  --volumes-from $OVPN_DATA \
  --rm \
  kylemanna/openvpn \
  ovpn_getclient $CLIENTNAME \
  > $CLIENTNAME.ovpn


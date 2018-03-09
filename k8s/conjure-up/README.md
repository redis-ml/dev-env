# Install Conjure-up

## Prepare instance.

1. Install lxd
```bash
sudo apt-get install snap
sudo snap install lxd
sudo /snap/bin/lxd init --auto
# This might be needed.
# sudo /snap/bin/lxc network create lxdbr0 ipv4.address=auto ipv4.nat=true ipv6.address=none ipv6.nat=false

## Install Conjure up
sudo snap install conjure-up --classic

## Bring up K8s
```
conjure-up kubernetes
```

## Solve problems

1. Remove accidentally created link:
```bash
sudo ip link set lxdbr0 down
# sudo apt install bridge-utils
sudo brctl delbr lxdbr0
```

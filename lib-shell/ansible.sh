#!/bin/bash

export ANSIBLE_INVENTORY="~/ansible/inventory/consul.py"

function ansible_inventory() {
  ansible-inventory -i $ANSIBLE_INVENTORY --list "$@"
}

#!/bin/bash

set -ex

ACCOUNT_ID=`aws sts get-caller-identity --output text --query 'Account'`

terraform apply \
   -var "account_id=${ACCOUNT_ID}" \
   -var 'region=us-west-2' \



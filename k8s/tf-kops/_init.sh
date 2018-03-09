terraform init \
  -backend-config="region=us-west-2" \
  -backend-config="bucket=usdcny-tfstate" \
  -backend-config="key=tf/kops.tfstate" \

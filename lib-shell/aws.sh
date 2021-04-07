alias aws_id='aws sts get-caller-identity'


function get_ec2_by_private_ip() {
  local ip=${1?private-ip-address}
  aws ec2 describe-instances \
    --filters Name=network-interface.addresses.private-ip-address,Values=$ip
}

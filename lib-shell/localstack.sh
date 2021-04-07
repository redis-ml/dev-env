
function localstack_aws() {
  local localstack_network="${LOCALSTACK_NETWORK:-localstack_localstack}"
  local localstack_endpoint="${LOCALSTACK_ENDPOINT:-http://localstack:4566}"
  docker run \
    --network "${localstack_network}" \
    --rm -it amazon/aws-cli \
    --endpoint-url="${localstack_endpoint}" \
    "$@"
}

function localstack_up() {
  docker-compose \
    -f "${MY_ENV_SCRIPT_DIR?}"/../docker-env/localstack/docker-compose.yaml \
    up
}

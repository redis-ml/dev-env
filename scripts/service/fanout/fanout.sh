#!/usr/bin/env bash

BASE_QUEUE_URL=https://us-west-2.queue.amazonaws.com/097605708335/fanout

QUEUE_URL="${BASE_QUEUE_URL}-input"

set -e
TS="$(date +%s)"
MSG="{\"start\":900000,\"end\":1600000,\"payload\":{\"po_box\":\"my_test\"},\"event_id\":\"event-${TS}\"}"

aws sqs send-message \
  --queue-url "${QUEUE_URL}" \
  --message-body "${MSG}"

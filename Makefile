
PATH := $(bazelisk run @go_sdk//:bin/go -- env GOROOT)/bin:$(PATH)

BAZEL := bazelisk

yarn:
	bazel run @yarn//:bin/yarn --

go:
	bazel run @go_sdk//:bin/go --

run_yq:
	go run github.com/mikefarah/yq/v4

# For fanout worker
build_fanout:
	$(BAZEL) run @go_sdk//:bin/go -- build github.com/redisliu/dev-env/golang/cmd/sqsfanoutworker/
	rm -f build/fanout/sqsfanoutworker.zip
	zip build/fanout/sqsfanoutworker.zip sqsfanoutworker

local_test_fanout:
	go run github.com/redisliu/dev-env/golang/cmd/sqsfanoutworker/localsqsfanout -filename data/examples/sqs/lambda-fanout-leaf.json

test_fanout:
	aws lambda invoke --function-name sqs-fanout-worker-dev --payload file://data/examples/sqs/lambda-fanout-leaf.json  --cli-binary-format raw-in-base64-out /dev/stdout

sqs_receive:
	# QUEUE_URL = https://us-west-2.queue.amazonaws.com/097605708335/fanout
	aws sqs receive-message --queue-url $(QUEUE_URL)

scan_fanout_ddb:
	# QUEUE_URL = https://us-west-2.queue.amazonaws.com/097605708335/fanout
	aws dynamodb scan --table-name "CommEvent"

provider "aws" {
  region  = "us-west-2"
}

data "aws_caller_identity" "current" {}

module "lambda_sqs_fanout" {
  source = "terraform-aws-modules/lambda/aws"

  for_each = toset(["dev"])

  provisioned_concurrent_executions = 10

  memory_size = 512

  timeout = 20

  function_name = "sqs-fanout-worker-${each.value}"
  description   = "SQS fanout worker"
  # The file name of the binary.
  handler       = "sqsfanoutworker"
  runtime       = "go1.x"
  publish       = true

  create_package         = false

  local_existing_package = "${path.root}/../../build/fanout/sqsfanoutworker.zip"

#   source_path = "${path.module}/../fixtures/python3.8-app1"

#  store_on_s3 = true
#  s3_bucket   = module.s3_bucket.s3_bucket_id
#  s3_prefix   = "lambda-builds/"

#  artifacts_dir = "${path.root}/.terraform/lambda-builds/"

  layers = [
#    module.lambda_layer_local.lambda_layer_arn,
#    module.lambda_layer_s3.lambda_layer_arn,
#    # this fails with /lib64/libc.so.6 GLIBC_2.18 not found.
#    "arn:aws:lambda:us-west-2:580247275435:layer:LambdaInsightsExtension:18",
    "arn:aws:lambda:us-west-2:580247275435:layer:LambdaInsightsExtension:14",
  ]

  environment_variables = {
    Hello      = "World"
    Serverless = "Terraform"
  }

  role_path   = "/tf-managed/"
  policy_path = "/tf-managed/"

  attach_dead_letter_policy = true
  dead_letter_target_arn    = aws_sqs_queue.dlq.arn

  attach_policy = true
  policy        = "arn:aws:iam::aws:policy/AWSXRayDaemonWriteAccess"

  attach_policies    = true
  policies           = [
    "arn:aws:iam::aws:policy/AWSXrayReadOnlyAccess",
    "arn:aws:iam::aws:policy/CloudWatchLambdaInsightsExecutionRolePolicy",
  ]
  number_of_policies = 2

  attach_policy_statements = true
  policy_statements = {
    dynamodb = {
      effect    = "Allow",
      actions   = [
        "dynamodb:DeleteItem",
        "dynamodb:GetItem",
        "dynamodb:PutItem",
        "dynamodb:UpdateItem",
        "dynamodb:Scan",
        "dynamodb:Query",
        "dynamodb:BatchWriteItem",
      ],
      resources = [aws_dynamodb_table.event.arn]
    },
    s3_list_bucket = {
      effect    = "Allow",
      actions   = [
        "s3:ListBucket", 
        "s3:ListBucketVersion", 
        "s3:GetBucketLocation",
      ],
      resources = [
        module.s3_bucket_for_fanout.s3_bucket_arn,
      ]
    },
    s3_rw = {
      effect    = "Allow",
      actions   = [
        "s3:DeleteObject*",
        "s3:PutObject*",
        "s3:GetObject*",
      ],
      resources = [
        "${module.s3_bucket_for_fanout.s3_bucket_arn}/*",
      ]
    },
    handle_fanout_input_sqs = {
      effect    = "Allow",
      actions = [
        "sqs:ReceiveMessage",
        "sqs:DeleteMessage",
        "sqs:GetQueueAttributes",
      ]
      resources = [
        aws_sqs_queue.fanout_input.arn,
      ]
    },
    fanout_sqs = {
      effect    = "Allow",
      actions = [
        "sqs:SendMessage",
        "sqs:ReceiveMessage",
        "sqs:DeleteMessage",
        "sqs:GetQueueAttributes",
      ]
      resources = [
        aws_sqs_queue.fanout.arn,
      ]
    },
    sink_sqs = {
      effect    = "Allow",
      actions = [
        "sqs:SendMessage",
        "sqs:GetQueueAttributes",
      ]
      resources = [
        aws_sqs_queue.fanout_sink.arn,
      ]
    },
  }
}

##################
# Extra resources
##################

resource "random_pet" "this" {
  length = 2
}

module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"

  bucket        = "redisliu-lambda-${random_pet.this.id}"
  force_destroy = true
}

resource "aws_sqs_queue" "dlq" {
  name = random_pet.this.id
}

###############################################################################################################
# Fanout used

resource "aws_sqs_queue" "fanout_input" {
  name = "fanout-input"
}

resource "aws_sqs_queue" "fanout" {
  name = "fanout"
}

################################################
# Store the data, mimicking the final consumer.

resource "aws_sqs_queue" "fanout_sink" {
  name = "fanout-sink"
}

module "s3_bucket_for_fanout" {
  source = "terraform-aws-modules/s3-bucket/aws"

  bucket        = "redisliu-fanout"
  force_destroy = true

  versioning = {
    status     = true
    mfa_delete = false
  }
}

resource "aws_dynamodb_table" "event" {
  name = "CommEvent"

  billing_mode = "PAY_PER_REQUEST"

  hash_key = "Owner"

  range_key = "EventID"

  attribute {
    name = "Owner"
    type = "S"
  }

  attribute {
    name = "EventID"
    type = "S"
  }

  attribute {
    name = "CreatedAt"
    type = "S"
  }

  ttl {
    attribute_name = "TTL"
    enabled        = true
  }

  local_secondary_index {
    name            = "IdxCreatedAt"
    range_key       = "CreatedAt"
    projection_type = "ALL"
  }

  server_side_encryption {
    enabled = true
  }

  point_in_time_recovery {
    enabled = true
  }
}

# Subscribe the SQS events to Lambda.
resource "aws_lambda_event_source_mapping" "fanout" {
  event_source_arn = aws_sqs_queue.fanout_input.arn

  enabled = true

  for_each = module.lambda_sqs_fanout
  function_name    = each.value.lambda_function_arn

  batch_size = 1
}

# Subscribe the SQS events to Lambda.
resource "aws_lambda_event_source_mapping" "fanout_spawn" {
  event_source_arn = aws_sqs_queue.fanout.arn

  enabled = true

  for_each = module.lambda_sqs_fanout
  function_name    = each.value.lambda_function_arn

  batch_size = 1
}

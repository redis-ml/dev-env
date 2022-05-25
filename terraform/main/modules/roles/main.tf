provider "aws" {
  region  = "us-west-2"
}

data "aws_iam_policy_document" "common" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "common" {
  name               = "common"
  assume_role_policy = data.aws_iam_policy_document.common.json
}

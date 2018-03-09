provider "aws" {
  version = "~> 1.7.0"
  region = "${var.region}"

  assume_role {
    role_arn = "arn:aws:iam::${var.account_id}:role/${var.tf_role}"
  }
}

terraform {
  backend "s3" {
    encrypt = "true"
  }
}

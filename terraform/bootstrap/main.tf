locals {
  region = "us-west-2"
}

terraform {
  required_version = "~> 1.1"
}

module "bootstrap" {
  # source = "github.com/redisliu/terraform-aws-bootstrap"
  source = "github.com/trussworks/terraform-aws-bootstrap"

  region        = local.region
  account_alias = "redis"
}

provider "aws" {
  version = "~> 3.0"
  region  = local.region
}

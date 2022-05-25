terraform {
  required_version = "~> 1.1"

  backend "s3" {
    bucket         = "redis-tf-state-us-west-2"
    key            = "dev-env/terraform.tfstate"
    dynamodb_table = "terraform-state-lock"
    region         = "us-west-2"
    encrypt        = "true"
  }

  required_providers {
    aws = {
      version = "~> 3.37.0"
    }
  }
}

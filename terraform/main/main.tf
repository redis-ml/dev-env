provider "aws" {
  alias   = "us_west_2"
  region  = "us-west-2"
}

provider "aws" {
  alias   = "us_east_1"
  region  = "us-east-1"
}

module "dns" {
  source = "./modules/dns"
}

module "roles" {
  source = "./modules/roles"
}

module "users" {
  source = "./modules/users"
}

module "dynamodb" {
  source = "./modules/dynamodb"
}

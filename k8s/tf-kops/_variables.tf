variable "tf_role" {
    default = "TF_PROVISIONER"
}
variable "account_id" {}
variable "region" {
    default = "us-west-2"
}

variable "tf_s3_bucket" {
    default = "usdcny-tfstate"
}
variable "tfstate_s3key" {
    default = "tf/ec2.tfstate"
}

variable "logging_bucket" {
    default = "usdcny-s3-access-log"
}
variable "account" {
    type = "string"
    default = "usdcny"
    description = "The AWS account a s3 bucket is associated with."
}
variable "scope" {
    type = "string"
    default = "internal-only"
    description = "The scope of the s3 bucket: internal-only, partner, or public."
}
variable "usage" {
    type = "string"
    default = "service"
    description = "The intended usage of the s3 bucket: human-user, service"
}
variable "internal_access" {
    type = "string"
    default = "read-write"
    description = "The accesses internal users have on a s3 bucket."
}
variable "partner_access" {
    type = "string"
    default = "none"
    description = "The accesses third-party partners have on a s3 bucket."
}
variable "public_access" {
    type = "string"
    default = "none"
    description = "The accesses public users have on a s3 bucket."
}
variable "purpose" {
    type = "string"
    default = ""
    description = "A short description of the intended purpose of the s3 bucket."
}

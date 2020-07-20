resource "aws_s3_bucket" "test1" {
  bucket = "usdcny-my-tf-test-bucket"
  acl    = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

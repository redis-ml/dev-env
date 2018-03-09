
resource "aws_s3_bucket" "clusters-dev-gods-tools" {
  bucket = "clusters.dev.gods.tools"
  acl    = "private"
}

resource "aws_route53_zone" "gods_tools" {
  name = "gods.tools"

  tags = {
    Source = "GoogleDomain"
  }
}

resource "aws_route53_zone" "dev_gods_tools" {
  name = "dev.gods.tools"

  tags = {
    Environment = "dev"
  }
}


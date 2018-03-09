resource "aws_route53_zone" "gods-tools" {
  name = "gods.tools"
}

resource "aws_route53_zone" "dev-gods-tools" {
  name = "dev.gods.tools"

  tags {
    Environment = "dev"
  }
}

resource "aws_route53_record" "dev-gods-tools" {
  zone_id = "${aws_route53_zone.gods-tools.zone_id}"
  name    = "dev.gods.tools"
  type    = "NS"
  ttl     = "30"

  records = [
    "${aws_route53_zone.dev-gods-tools.name_servers.0}",
    "${aws_route53_zone.dev-gods-tools.name_servers.1}",
    "${aws_route53_zone.dev-gods-tools.name_servers.2}",
    "${aws_route53_zone.dev-gods-tools.name_servers.3}",
  ]
}

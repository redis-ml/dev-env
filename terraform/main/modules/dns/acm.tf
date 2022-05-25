resource "aws_acm_certificate" "wildcard_gods_tools" {
  domain_name       = "*.gods.tools"
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "acm_verify_wildcard_gods_tools" {
  for_each = {
    for dvo in aws_acm_certificate.wildcard_gods_tools.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type

  zone_id = aws_route53_zone.gods_tools.id
}

resource "aws_acm_certificate_validation" "wildcard_gods_tools" {
  certificate_arn         = aws_acm_certificate.wildcard_gods_tools.arn
  validation_record_fqdns = [for record in aws_route53_record.acm_verify_wildcard_gods_tools : record.fqdn]
}

# resource "aws_acm_certificate" "wildcard_dev_gods_tools" {
#   domain_name       = "*.dev.gods.tools"
#   validation_method = "DNS"
# 
#   lifecycle {
#     create_before_destroy = true
#   }
# }
# 
# resource "aws_route53_record" "acm_verify_wildcard_dev_gods_tools" {
#   for_each = {
#     for dvo in aws_acm_certificate.wildcard_dev_gods_tools.domain_validation_options : dvo.domain_name => {
#       name   = dvo.resource_record_name
#       record = dvo.resource_record_value
#       type   = dvo.resource_record_type
#     }
#   }
# 
#   allow_overwrite = true
#   name            = each.value.name
#   records         = [each.value.record]
#   ttl             = 60
#   type            = each.value.type
# 
#   zone_id = aws_route53_zone.dev_gods_tools.id
# }
# 
# resource "aws_acm_certificate_validation" "wildcard_dev_gods_tools" {
#   certificate_arn         = aws_acm_certificate.wildcard_dev_gods_tools.arn
#   validation_record_fqdns = [for record in aws_route53_record.acm_verify_wildcard_dev_gods_tools : record.fqdn]
# }
# 
# ##########################################################################################
# 
# resource "aws_acm_certificate" "sg_dev_gods_tools" {
#   domain_name       = "sg.dev.gods.tools"
#   validation_method = "DNS"
# 
#   lifecycle {
#     create_before_destroy = true
#   }
# }
# 
# resource "aws_route53_record" "acmverify_sg_dev_gods_tools" {
#   for_each = {
#     for dvo in aws_acm_certificate.sg_dev_gods_tools.domain_validation_options : dvo.domain_name => {
#       name   = dvo.resource_record_name
#       record = dvo.resource_record_value
#       type   = dvo.resource_record_type
#     }
#   }
# 
#   allow_overwrite = true
#   name            = each.value.name
#   records         = [each.value.record]
#   ttl             = 60
#   type            = each.value.type
# 
#   zone_id = aws_route53_zone.dev_gods_tools.id
# }
# 
# resource "aws_acm_certificate_validation" "sg_dev_gods_tools" {
#   certificate_arn         = aws_acm_certificate.sg_dev_gods_tools.arn
#   validation_record_fqdns = [for record in aws_route53_record.acmverify_sg_dev_gods_tools : record.fqdn]
# }
# 
# # Mimic the case when the dns was already set.
# resource "aws_route53_record" "sg_dev_gods_tools" {
#   allow_overwrite = true
# 
#   name    = "sg.dev.gods.tools"
#   type    = "CNAME"
#   ttl     = 60
# 
#   records        = [aws_cloudfront_distribution.sg_dev_gods_tools.domain_name]
# 
#   zone_id = aws_route53_zone.dev_gods_tools.id
# }
# 
# ##########################################################################################
# 
# resource "aws_cloudfront_distribution" "sg_dev_gods_tools" {
#   enabled         = true
#   is_ipv6_enabled = true
# 
#   #aliases = ["*.sg.dev.gods.tools", "x.sg.dev.gods.tools"]
#   aliases = ["sg.dev.gods.tools"]
# 
#   origin {
#     domain_name = "sendgrid.net"
#     origin_id   = "sendgrid_net"
# 
#     custom_origin_config {
#       http_port = 80
#       https_port = 443
# 
#       origin_protocol_policy = "https-only"
#       origin_ssl_protocols = ["SSLv3"]
#     }
#   }
# 
#   restrictions {
#     geo_restriction {
#       restriction_type = "none"
#     }
#   }
# 
#   viewer_certificate {
#     # acm_certificate_arn      = aws_acm_certificate.wildcard_sg_dev_gods_tools.arn
#     acm_certificate_arn      = aws_acm_certificate.sg_dev_gods_tools.arn
#     ssl_support_method       = "sni-only"
#     minimum_protocol_version = "TLSv1.1_2016"
#   }
# 
#   default_cache_behavior {
#     allowed_methods  = ["GET", "HEAD"]
#     cached_methods   = ["GET", "HEAD"]
#     target_origin_id = "sendgrid_net"
# 
#     forwarded_values {
#       query_string = true
# 
#       cookies {
#         forward = "none"
#       }
#     }
# 
#     # This is necessary.
#     trusted_signers = []
# 
#     # viewer_protocol_policy = "redirect-to-https"
#     viewer_protocol_policy = "allow-all"
#     max_ttl                = 86400
#   }
# }
# 
# # # Add the resource to AWS Shield for DDoS protection
# # resource "aws_shield_protection" "cloudfront_sg_dev_gods_tools" {
# #   name         = aws_cloudfront_distribution.sg_dev_gods_tools.domain_name
# #   resource_arn = aws_cloudfront_distribution.sg_dev_gods_tools.arn
# # }

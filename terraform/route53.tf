resource "aws_route53_zone" "adventar_org" {
  name    = "adventar.org."
  comment = "Hosted Zone for adventar.org"
}

resource "aws_route53_record" "adventar_org" {
  zone_id = aws_route53_zone.adventar_org.id
  name    = "adventar.org"
  type    = "A"
  records = ["139.59.194.254"]
  ttl     = 300
}

resource "aws_route53_record" "www_adventar_org" {
  zone_id = aws_route53_zone.adventar_org.id
  name    = "www.adventar.org"
  type    = "CNAME"
  records = [aws_cloudfront_distribution.www_adventar_org.domain_name]
  ttl     = 300
}

resource "aws_route53_record" "dev_adventar_org" {
  zone_id = aws_route53_zone.adventar_org.id
  name    = "dev.adventar.org"
  type    = "A"

  alias {
    name    = aws_cloudfront_distribution.main.domain_name
    zone_id = aws_cloudfront_distribution.main.hosted_zone_id

    evaluate_target_health = true
  }
}

resource "aws_route53_record" "api_adventar_org" {
  zone_id = aws_route53_zone.adventar_org.id
  name    = "api.adventar.org"
  type    = "A"

  alias {
    name    = aws_alb.api.dns_name
    zone_id = aws_alb.api.zone_id

    evaluate_target_health = true
  }
}

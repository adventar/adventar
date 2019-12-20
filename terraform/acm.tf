resource "aws_acm_certificate" "adventar_org" {
  provider    = aws.us-east-1
  domain_name = "adventar.org"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate" "www_adventar_org" {
  provider    = aws.us-east-1
  domain_name = "www.adventar.org"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate" "api_adventar_org" {
  domain_name = "api.adventar.org"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate" "img_adventar_org" {
  provider    = aws.us-east-1
  domain_name = "img.adventar.org"

  lifecycle {
    create_before_destroy = true
  }
}

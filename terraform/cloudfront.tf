locals {
  s3_adventar_assets_origin_id = "s3-adventar-assets"
  frontend_server_origin_id    = "frontend-server"
}

resource "aws_cloudfront_origin_access_identity" "s3_adventar_assets" {}

resource "aws_cloudfront_distribution" "main" {
  aliases = ["adventar.org"]

  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"

  origin {
    origin_id   = local.s3_adventar_assets_origin_id
    domain_name = aws_s3_bucket.adventar_assets.bucket_regional_domain_name
    origin_path = "/nuxt"

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.s3_adventar_assets.cloudfront_access_identity_path
    }
  }

  origin {
    origin_id   = local.frontend_server_origin_id
    domain_name = "2z1t0nefxk.execute-api.ap-northeast-1.amazonaws.com"
    origin_path = "/prod"

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols   = ["TLSv1.2"]
    }
  }

  custom_error_response {
    error_code         = "403"
    response_code      = "200"
    response_page_path = "/index.html"
  }

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = local.s3_adventar_assets_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 31536000
    max_ttl                = 31536000
    compress               = true
  }

  ordered_cache_behavior {
    path_pattern     = "/calendars/*"
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = local.frontend_server_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 0
    max_ttl                = 31536000
    compress               = true
  }

  ordered_cache_behavior {
    path_pattern     = "/users/*.ics"
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = local.frontend_server_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 0
    max_ttl                = 31536000
    compress               = true
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.adventar_org.arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1"
  }
}

locals {
  www_adventar_org_origin_id = "www-adventar-org"
}

resource "aws_cloudfront_distribution" "www_adventar_org" {
  enabled = true
  aliases = ["www.adventar.org"]

  origin {
    origin_id   = local.www_adventar_org_origin_id
    domain_name = aws_s3_bucket.www_adventar_org.website_endpoint

    custom_origin_config {
      origin_protocol_policy = "http-only"
      http_port              = "80"
      https_port             = "443"
      origin_ssl_protocols   = ["TLSv1"]
    }
  }

  default_cache_behavior {
    target_origin_id       = local.www_adventar_org_origin_id
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    viewer_protocol_policy = "allow-all"

    forwarded_values {
      query_string = true

      cookies {
        forward = "none"
      }
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.www_adventar_org.arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1"
  }
}

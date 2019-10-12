locals {
  s3_adventar_assets_origin_id = "s3-adventar-assets"
}

resource "aws_cloudfront_origin_access_identity" "s3_adventar_assets" {}

resource "aws_cloudfront_distribution" "main" {
  aliases = ["dev.adventar.org"]

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
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.dev_adventar_org.arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1"
  }
}

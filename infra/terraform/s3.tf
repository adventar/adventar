resource "aws_s3_bucket" "adventar" {
  bucket = "adventar"
  acl    = "private"
}

data "aws_elb_service_account" "alb_log" {}

data "aws_iam_policy_document" "s3_adventar_policy" {
  statement {
    actions   = ["s3:PutObject"]
    resources = ["${aws_s3_bucket.adventar.arn}/alb-logs/*"]

    principals {
      type        = "AWS"
      identifiers = [data.aws_elb_service_account.alb_log.id]
    }
  }
}

resource "aws_s3_bucket_policy" "adventar" {
  bucket = aws_s3_bucket.adventar.id
  policy = data.aws_iam_policy_document.s3_adventar_policy.json
}

resource "aws_s3_bucket" "adventar_assets" {
  bucket = "adventar-assets"
  acl    = "private"
}

data "aws_iam_policy_document" "s3_adventar_assets_policy" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.adventar_assets.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.s3_adventar_assets.iam_arn]
    }
  }
}

resource "aws_s3_bucket_policy" "adventar_assets" {
  bucket = aws_s3_bucket.adventar_assets.id
  policy = data.aws_iam_policy_document.s3_adventar_assets_policy.json
}

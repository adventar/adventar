resource "aws_s3_bucket" "adventar" {
  bucket = "adventar"
  acl    = "private"
  region = "ap-northeast-1"
}

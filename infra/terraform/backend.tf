terraform {
  backend "s3" {
    bucket = "adventar"
    key    = "terraform/adventar.tfstate"
    region = "ap-northeast-1"
  }
}

provider "aws" {
  version             = "= 2.29.0"
  region              = "ap-northeast-1"
  allowed_account_ids = ["287379415997"]
}

provider "aws" {
  alias               = "us-east-1"
  version             = "= 2.29.0"
  region              = "us-east-1"
  allowed_account_ids = ["287379415997"]
}

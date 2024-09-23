terraform {
  backend "s3" {
    region = "ap-northeast-1"
    # ローカルの設定 start
    use_path_style = true
    endpoints = {
      s3  = "http://127.0.0.1:5000"
      sts = "http://127.0.0.1:5000"
    }
    # ローカルの設定 end
  }
  required_version = "1.8.3"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.55"
    }
  }
}

provider "aws" {
  region = var.region
  default_tags {
    tags = {
      Service = var.service
      Env     = var.env
    }
  }
  # ローカルの設定 start
  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  endpoints {
    s3       = "http://127.0.0.1:5000"
    ec2      = "http://127.0.0.1:5000"
    dynamodb = "http://localhost:8000"
  }
  # ローカルの設定 end
}

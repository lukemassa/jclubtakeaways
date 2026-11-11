terraform {
  backend "s3" {}

  required_version = "~> 1.17.1"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.72"
    }
  }
}

provider "aws" {
  region  = "us-east-1"
  profile = "AdministratorAccess-917497682277"
}

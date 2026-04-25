locals {
  aws_region = "us-east-1"
}

remote_state {
  backend = "s3"

  config = {
    bucket         = "tg-state-${get_aws_account_id()}"
    key            = "terraform.tfstate"
    region         = local.aws_region
    encrypt        = true
    dynamodb_table = "tg-locks"
  }
}

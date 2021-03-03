terraform {
  backend "s3" {
    bucket         = "seanjh-117936299034-terraform-state"
    dynamodb_table = "seanjh-terraform-locks"
    encrypt        = true
    key            = "dev/terraform.tfstate"
    region         = "us-east-1"
  }
}

provider "aws" {
  region = "us-east-1"
  assume_role {
    role_arn = "arn:aws:iam::713237909615:role/admin"
  }
}

module "serverless" {
  source = "../modules/serverless"
}

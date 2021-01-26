terraform {
  required_version = "0.14.4"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.24.0"
    }
  }

  # via https://www.terraform.io/docs/backends/types/s3.html
  backend "s3" {
    bucket         = "seanjh-117936299034-terraform-state"
    dynamodb_table = "seanjh-terraform-locks"
    encrypt        = true
    key            = "global/s3/terraform.tfstate"
    region         = "us-east-1"
  }
}

provider "aws" {
  region = "us-east-1"
}

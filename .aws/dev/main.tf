terraform {
  # via https://www.terraform.io/docs/backends/types/s3.html
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
}

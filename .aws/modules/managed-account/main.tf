provider "aws" {
  alias  = "parent_account"
  region = var.region
}

provider "aws" {
  alias  = "managed_account"
  region = var.region

  assume_role {
    role_arn = "arn:aws:iam::${aws_organizations_account.managed_account.id}:role/OrganizationAccountAccessRole"
  }
}

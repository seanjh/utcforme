module "aws_account_dev" {
  source = "../modules/managed-account"

  name  = "dev"
  email = "seanherman+dev@gmail.com"
}

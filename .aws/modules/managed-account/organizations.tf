resource "aws_organizations_account" "managed_account" {
  provider = aws.parent_account

  name  = var.name
  email = var.email

  iam_user_access_to_billing = "DENY"
}

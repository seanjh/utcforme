output "managed_account" {
  value = aws_organizations_account.managed_account
}

output "iam_roles" {
  value = {
    admin = aws_iam_role.admin
  }
}

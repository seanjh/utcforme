output "s3_bucket_arn" {
  value       = aws_s3_bucket.terraform_state.arn
  description = "The ARN of the Terraform state S3 bucket"
}

output "dynamodb_table_name" {
  value       = aws_dynamodb_table.terraform_locks.name
  description = "The name of the Terraform locks DynamoDB table"
}

output "aws_organization" {
  description = "AWS organization"
  value       = aws_organizations_organization.org
}

output "aws_managed_account_dev" {
  description = "AWS organization account - dev"
  value       = module.aws_account_dev.managed_account
}

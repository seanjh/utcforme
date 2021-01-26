variable "name" {
  description = "AWS account name (required)."
  type        = string
}

variable "email" {
  description = "AWS account root user email address (required)."
  type        = string
}

variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

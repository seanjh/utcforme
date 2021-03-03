data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda_exec" {
  name        = "serverless_lambda"
  description = "Lambda execution role"

  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

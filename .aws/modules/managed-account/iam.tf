#resource "aws_iam_account_alias" "alias" {
#provider = aws.managed_account

#account_alias = var.name
#}

data "aws_caller_identity" "parent_account" {
  provider = aws.parent_account
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.parent_account.account_id}:root"
      ]
    }
  }
}

resource "aws_iam_role" "admin" {
  provider = aws.managed_account

  name        = "admin"
  description = "Administrative access"

  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy_attachment" "admin_access" {
  provider = aws.managed_account

  role       = aws_iam_role.admin.name
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}

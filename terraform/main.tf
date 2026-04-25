resource "aws_iam_role" "token_minter" {
  name = "token-minter"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_ssm_parameter" "webclientaccount_key" {
  name  = "/token-minter/webclientaccount_key"
  type  = "SecureString"
  value = "placeholder"

  lifecycle {
    ignore_changes = [value]
  }
}

resource "aws_iam_role_policy" "webclientaccount_key_access" {
  role = aws_iam_role.token_minter.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Action = [
        "ssm:GetParameter"
      ]
      Resource = aws_ssm_parameter.webclientaccount_key.arn
    }]
  })
}

resource "aws_iam_role_policy_attachment" "token_minter_attachment" {
  role       = aws_iam_role.token_minter.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "token_minter" {
  function_name = "token-minter"

  role    = aws_iam_role.token_minter.arn
  runtime = "provided.al2023"
  handler = "bootstrap"

  architectures = ["arm64"]

  filename         = "placeholder.zip"
  source_code_hash = filebase64sha256("placeholder.zip")

  lifecycle {
    ignore_changes = [
      filename,
      source_code_hash,
    ]
  }
}

moved {
  from = aws_lambda_function.token-minter
  to = aws_lambda_function.token_minter
}

resource "aws_lambda_function_url" "token_minter" {
  function_name      = aws_lambda_function.token_minter.function_name
  authorization_type = "NONE"
}

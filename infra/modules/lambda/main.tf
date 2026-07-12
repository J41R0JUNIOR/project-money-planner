data "archive_file" "this" {
  type        = "zip"
  source_file = var.source_file_path
  output_path = "${path.module}/function.zip"
}

resource "aws_lambda_function" "this" {
  filename         = data.archive_file.this.output_path
  function_name    = var.function_name
  role             = var.role_arn
  handler          = "index.handler"
  source_code_hash = data.archive_file.this.output_base64sha256
  runtime          = "go1.x"

  environment {
    variables = merge(
      {
        APP_NAME    = var.app_name
        ENVIRONMENT = var.environment
      },
      var.env_vars
    )
  }
}
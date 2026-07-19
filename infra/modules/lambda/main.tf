# data "archive_file" "this" {
#   type        = "zip"
#   source_file = var.zip_path
#   output_path = "${path.module}/function.zip"
# }


resource "aws_lambda_function" "this" {
  filename         = var.zip_path
  function_name    = var.function_name
  role             = var.role_arn
  handler          = "bootstrap"
  runtime          = "provided.al2023"

  architectures = ["arm64"]

  source_code_hash = filebase64sha256(var.zip_path)

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
# Provider configuration
provider "aws" {
  region = var.aws_region
  access_key = var.access_key
  secret_key = var.secret_key
}

# Resources
resource "aws_cognito_user_pool" "user_pool" {
  name                     = "user_pool_${var.environment}"
  auto_verified_attributes = ["email"]
}

resource "aws_cognito_user_pool_client" "user_pool_client" {
    name = "${var.app_name}_client_${var.environment}"

    user_pool_id = aws_cognito_user_pool.user_pool.id
}

# DynamoDb Resource
# Preciso entender melhor como vai ser o schema do banco anntes
# resource "aws_dynamodb_table" "dynamo_table" {
#     name = "${var.app_name}_table_${var.environment}"
#     billing_mode = "PAY_PER_REQUEST"
#     read_capacity = 20
#     write_capacity = 20
#     hash_key = "id"
# }

# Lambda Functions
# resource "aws_lambda_function" "auth_lambda" {}
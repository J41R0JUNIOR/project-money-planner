resource "aws_cognito_user_pool" "this" {
  name                     = "${var.app_name}-user-pool-${var.environment}"
  auto_verified_attributes = ["email"]
}

resource "aws_cognito_user_pool_client" "this" {
  name         = "${var.app_name}-user-pool-client-${var.environment}"
  user_pool_id = aws_cognito_user_pool.this.id
}
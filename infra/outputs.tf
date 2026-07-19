output "cognito_user_pool_id" {
  value = module.cognito.user_pool_id
}

output "cognito_user_pool_client_id" {
  value = module.cognito.user_pool_client_id
}

output "lambda_arns" {
  value = { for name, lambda in module.lambda : name => lambda.lambda_arn }
}

output "lambda_role_arn" {
  value = module.iam.lambda_role_arn
}

output "api_url" {
    value = module.apigateway.api_endpoint
}


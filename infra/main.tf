locals {
  name_prefix = lower(replace(var.app_name, " ", "-"))
  lambda_path = "internal/lambdas"

  lambdas = {
    auth = {
      function_name    = "${local.name_prefix}-${var.environment}-auth"
      source_file_path = "${local.lambda_path}/auth/index.go"
      env_vars = {
        FUNCTION_TYPE = "auth"
      }
    }

    user = {
      function_name    = "${local.name_prefix}-${var.environment}-user"
      source_file_path = "${local.lambda_path}/user/index.go"
      env_vars = {
        FUNCTION_TYPE = "user"
      }
    }

    transaction = {
      function_name    = "${local.name_prefix}-${var.environment}-transaction"
      source_file_path = "${local.lambda_path}/transaction/index.go"
      env_vars = {
        FUNCTION_TYPE = "transaction"
      }
    }
  }
}

module "cognito" {
  source      = "./modules/cognito"
  app_name    = var.app_name
  environment = var.environment
}

module "iam" {
  source      = "./modules/iam"
  app_name    = var.app_name
  environment = var.environment
}

module "lambda" {
  for_each = local.lambdas

  source           = "./modules/lambda"
  app_name         = var.app_name
  environment      = var.environment
  function_name    = each.value.function_name
  role_arn         = module.iam.lambda_role_arn
  source_file_path = each.value.source_file_path
  env_vars         = each.value.env_vars
}
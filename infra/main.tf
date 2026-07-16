locals {
  app_name = lower(replace(var.app_name, " ", "-"))

  lambdas = {
    auth = {
      function_name = "${local.app_name}-${var.environment}-auth"
      zip_path      = abspath("${path.root}/../dist/auth/function.zip")
      env_vars = {
        COGNITO_USER_POOL_ID = module.cognito.user_pool_id
        COGNITO_CLIENT_ID    = module.cognito.user_pool_client_id
      }
    }

    account = {
      function_name = "${local.app_name}-${var.environment}-account"
      zip_path      = abspath("${path.root}/../dist/account/function.zip")
    }
  }
}

module "cognito" {
  source      = "./modules/cognito"
  app_name    = local.app_name
  environment = var.environment
}

module "iam" {
  source      = "./modules/iam"
  app_name    = local.app_name
  environment = var.environment
}

module "dynamodb" {
  source      = "./modules/dynamodb"
}

module "lambda" {
  for_each = local.lambdas

  source        = "./modules/lambda"
  app_name      = local.app_name
  environment   = var.environment
  function_name = each.value.function_name
  role_arn      = module.iam.lambda_role_arn
  zip_path      = each.value.zip_path

  env_vars = lookup(each.value, "env_vars", {})
}
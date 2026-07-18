locals {
  app_name = lower(replace(var.app_name, " ", "-"))

  lambdas = {
    auth = {
      function_name = "${local.app_name}-${var.environment}-auth"
      zip_path      = abspath("${path.root}/../dist/auth/function.zip")
      public        = true
      env_vars = {
        COGNITO_CLIENT_ID    = module.cognito.user_pool_client_id
        COGNITO_USER_POOL_ID = module.cognito.user_pool_id

        DYNAMODB_TABLE_NAME = module.dynamodb.table_name
      }
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
  app_name    = local.app_name
  environment = var.environment
}

module "lambda" {
  for_each = local.lambdas

  source        = "./modules/lambda"
  app_name      = local.app_name
  environment   = var.environment
  function_name = each.value.function_name
  role_arn      = module.iam.lambda_role_arn
  zip_path      = each.value.zip_path

  public   = each.value.public
  env_vars = lookup(each.value, "env_vars", {})
}

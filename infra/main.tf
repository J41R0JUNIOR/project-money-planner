locals {
  app_name = lower(replace(var.app_name, " ", "-"))

  lambdas = {
    auth = {
      function_name = "${local.app_name}-${var.environment}-auth"
      zip_path      = abspath("${path.root}/../dist/auth/function.zip")
      public        = true
      routes = [
        {
          method = "POST"
          path   = "/auth/signup"
        },
        {
          method = "POST"
          path   = "/auth/confirm-code"
        },
        {
          method = "POST"
          path   = "/auth/signin"
        },
        {
          method = "POST"
          path   = "/auth/refresh"
        }
      ]
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

  env_vars = lookup(each.value, "env_vars", {})
}

module "apigateway" {
  source = "./modules/apigateway"

  app_name = local.app_name

  cognito_user_pool_id = module.cognito.user_pool_id
  cognito_client_id    = module.cognito.user_pool_client_id

  lambdas = {
    for key, value in local.lambdas :

    key => {
      lambda_name = module.lambda[key].lambda_name
      lambda_arn  = module.lambda[key].lambda_arn

      public = value.public
      routes = value.routes
    }
  }
}

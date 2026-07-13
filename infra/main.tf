locals {
  app_name = lower(replace(var.app_name, " ", "-"))

  lambdas = {
    auth = {
      function_name = "${local.app_name}-${var.environment}-auth"
      zip_path      = abspath("${path.root}/../dist/auth/function.zip")
    }

    user = {
      function_name = "${local.app_name}-${var.environment}-user"
      zip_path      = abspath("${path.root}/../dist/user/function.zip")
    }

    # transaction = {
    #   function_name = "${local.app_name}-${var.environment}-transaction"
    #   zip_path      = abspath("${path.root}/../dist/transaction/function.zip")
    # }
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

module "lambda" {
  for_each = local.lambdas

  source        = "./modules/lambda"
  app_name      = local.app_name
  environment   = var.environment
  function_name = each.value.function_name
  role_arn      = module.iam.lambda_role_arn
  zip_path      = each.value.zip_path

  # Descomente quando adicionar variáveis de ambiente específicas
  # env_vars = lookup(each.value, "env_vars", {})
}
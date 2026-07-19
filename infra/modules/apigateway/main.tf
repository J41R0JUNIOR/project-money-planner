resource "aws_apigatewayv2_api" "this" {

  name = "${var.app_name}-api"

  protocol_type = "HTTP"

  cors_configuration {

    allow_origins = [
      "http://localhost:5173"
    ]

    allow_methods = ["*"]

    allow_headers = ["*"]
  }
}

resource "aws_apigatewayv2_stage" "default" {

  api_id = aws_apigatewayv2_api.this.id

  name = "$default"

  auto_deploy = true
}

data "aws_region" "current" {}

resource "aws_apigatewayv2_authorizer" "cognito" {

  api_id = aws_apigatewayv2_api.this.id

  name = "cognito"

  authorizer_type = "JWT"

  identity_sources = [
    "$request.header.Authorization"
  ]

  jwt_configuration {

    audience = [
      var.cognito_client_id
    ]

    issuer = "https://cognito-idp.${data.aws_region.current.region}.amazonaws.com/${var.cognito_user_pool_id}"
  }
}

locals {

  routes = flatten([
    for lambda_name, lambda in var.lambdas : [
      for route in lambda.routes : {

        key = "${lambda_name}-${route.method}-${replace(route.path, "/", "-")}"

        lambda_name = lambda_name

        method = route.method
        path   = route.path

        public = lambda.public

        lambda_arn = lambda.lambda_arn
      }
    ]
  ])

  routes_map = {

    for r in local.routes :

    r.key => r
  }
}

resource "aws_apigatewayv2_integration" "lambda" {

  for_each = var.lambdas

  api_id = aws_apigatewayv2_api.this.id

  integration_type = "AWS_PROXY"

  integration_uri = each.value.lambda_arn

  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "route" {

  for_each = local.routes_map

  api_id = aws_apigatewayv2_api.this.id

  route_key = "${each.value.method} ${each.value.path}"

  target = "integrations/${aws_apigatewayv2_integration.lambda[split("-", each.key)[0]].id}"

  authorization_type = each.value.public ? "NONE" : "JWT"

  authorizer_id = each.value.public ? null : aws_apigatewayv2_authorizer.cognito.id
}

resource "aws_lambda_permission" "gateway" {

  for_each = var.lambdas

  statement_id = "AllowGateway-${each.key}"

  action = "lambda:InvokeFunction"

  function_name = each.value.lambda_name

  principal = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.this.execution_arn}/*"
}
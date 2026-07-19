variable "app_name" {
  type = string
}

variable "cognito_user_pool_id" {
  type = string
}

variable "cognito_client_id" {
  type = string
}

variable "lambdas" {
  type = map(object({

    lambda_name = string
    lambda_arn  = string

    public = bool

    routes = list(object({
      method = string
      path   = string
    }))
  }))
}
variable "app_name" {
  type        = string
  description = "The name of the application"
}

variable "environment" {
  type        = string
  description = "Deployment environment name"
}

variable "function_name" {
  type        = string
  description = "Lambda function name"
}

variable "role_arn" {
  type        = string
  description = "IAM role ARN for the Lambda execution role"
}

variable "source_file_path" {
  type        = string
  description = "Path to the Lambda source file"
}

variable "env_vars" {
  type        = map(string)
  description = "Additional environment variables for the Lambda"
  default     = {}
}
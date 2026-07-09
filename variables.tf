variable "aws_region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "us-east-2"

  validation {
    condition     = length(var.aws_region) > 0
    error_message = "aws_region must not be empty"
  }
}

variable "access_key" {
    type        = string
    description = "AWS Access Key"
}

variable "secret_key" {
    type        = string
    description = "AWS Secret Key"
}

variable "app_name" {
  type        = string
  description = "The name of the application"
  default     = "Money Planner"
}

variable "environment" {
  type        = string
  description = "Deployment environment name"
  default     = "dev"

  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Environment must be one of: dev, staging, prod"
  }
}
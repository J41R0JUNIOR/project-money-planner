locals {
  cors_origins = [
    data.terraform_remote_state.frontend.outputs.frontend_url
  ]
}

data "terraform_remote_state" "frontend" {

  backend = "s3"

  config = {
    bucket = "money-planner-tfstate-124355634510"
    key    = "frontend/dev/terraform.tfstate"
    region = "us-east-1"
  }
}
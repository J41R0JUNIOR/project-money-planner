resource "aws_dynamodb_table" "app_table" {
  name         = "${var.app_name}-${var.environment}"
  billing_mode = "PAY_PER_REQUEST"

  hash_key  = "PK"
  range_key = "SK"

  attribute {
    name = "PK"
    type = "S"
  }

  attribute {
    name = "SK"
    type = "S"
  }

  attribute {
    name = "GSI1PK"
    type = "S"
  }

  attribute {
    name = "GSI1SK"
    type = "S"
  }

  global_secondary_index {
    name = "GSI1"
    key_schema {
      attribute_name = "GSI1PK"
      key_type       = "HASH"
    }
    
    key_schema {
      attribute_name = "GSI1SK"
      key_type       = "RANGE"
    }

    projection_type = "ALL"
  }

  tags = {
    Application = var.app_name
    Environment = var.environment
  }
}

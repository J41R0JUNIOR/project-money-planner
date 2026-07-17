resource "aws_dynamodb_table" "app_table" {
    attribute {
        name = "id"
        type = "S"
    }

    hash_key = "id"
    name               = "${var.app_name}-lambda-role-${var.environment}"
    billing_mode = "PROVISIONED"
    read_capacity  = 5
    write_capacity = 5
}
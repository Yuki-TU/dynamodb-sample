resource "aws_dynamodb_table" "users_table" {
  name           = "Users"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "UserId"
  attribute {
    name = "UserId"
    type = "N"
  }
  attribute {
    name = "FirstName"
    type = "S"
  }
  attribute {
    name = "FirstNameKana"
    type = "S"
  }
  attribute {
    name = "LastName"
    type = "S"
  }
  attribute {
    name = "LastNameKana"
    type = "S"
  }
  ttl {
    attribute_name = "TimeToExist"
    enabled        = true
  }
  global_secondary_index {
    name            = "FirstNameIndex"
    hash_key        = "FirstName"
    range_key       = "UserId"
    write_capacity  = 20
    read_capacity   = 20
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "LastNameIndex"
    hash_key        = "LastName"
    range_key       = "UserId"
    write_capacity  = 20
    read_capacity   = 20
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "LastNameKanaIndex"
    hash_key        = "LastNameKana"
    range_key       = "UserId"
    write_capacity  = 20
    read_capacity   = 20
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "FirstNameKanaIndex"
    hash_key        = "FirstNameKana"
    range_key       = "UserId"
    write_capacity  = 20
    read_capacity   = 20
    projection_type = "ALL"
  }
}

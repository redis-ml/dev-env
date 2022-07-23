################################################
# Store the data, mimicking the final consumer.

# Schema and example:
# Owner: 0001
# PrefType: stock#111#custom_price_threshold#100.0
# InstrPref: stock#111#custom_price_threshold
# Thld: 100.0

resource "aws_dynamodb_table" "notification_preference" {
  name = "NotifPref"

  billing_mode = "PAY_PER_REQUEST"

  hash_key = "Owner"

  range_key = "PrefType"

  attribute {
    name = "Owner"
    type = "S"
  }

  attribute {
    name = "PrefType"
    type = "S"
  }

  attribute {
    name = "InstrPref"
    type = "S"
  }

  attribute {
    name = "Thld"
    type = "N"
  }

  ttl {
    attribute_name = "TTL"
    enabled        = true
  }

  global_secondary_index {
    name      = "IdxInstrPref"
    hash_key  = "InstrPref"

    projection_type = "KEYS_ONLY"
  }

  global_secondary_index {
    name      = "IdxInstrThldPref"
    hash_key  = "InstrPref"
    range_key = "Thld"

    projection_type = "KEYS_ONLY"
  }

  server_side_encryption {
    enabled = true
  }

  point_in_time_recovery {
    enabled = true
  }
}

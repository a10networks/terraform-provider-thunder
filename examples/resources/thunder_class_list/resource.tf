provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_class_list" "thunder_class_list" {
  name = "test"
  ac_list {
    ac_match_type = "contains"
    ac_key_string = "155"
    ac_value      = "87"
  }

  type     = "ac"
  user_tag = "14"
}
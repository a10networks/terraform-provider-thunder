provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_extcommunity_list_expanded_num" "thunder_ip_extcommunity_list_expanded_num" {

  ext_list_num = 129
  rules_list {
    ext_list_action = "deny"
    ext_list_value  = "test"
  }
}
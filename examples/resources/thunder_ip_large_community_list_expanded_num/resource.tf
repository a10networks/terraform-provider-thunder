provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_large_community_list_expanded_num" "thunder_ip_large_community_list_expanded_num" {

  ext_list_num = 100
  rules_list {
    ext_list_lcom_action = "deny"
    ext_list_lcom_value  = "test"
  }
}
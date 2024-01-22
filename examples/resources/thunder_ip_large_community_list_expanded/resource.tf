provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_large_community_list_expanded" "thunder_ip_large_community_list_expanded" {
  expanded = "7"
  rules_list {
    expanded_lcom_action = "deny"
  }
}
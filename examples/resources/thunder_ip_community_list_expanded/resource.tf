provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_community_list_expanded" "thunder_ip_community_list_expanded" {
  expanded = "test"
  rules_list {
    expanded_action = "deny"
    expanded_value  = "internet"
  }
}
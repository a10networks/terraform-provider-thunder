provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_community_list_standard" "thunder_ip_community_list_standard" {
  standard = "test"
  rules_list {
    standard_action     = "deny"
    standard_comm_value = "internet"
  }

}
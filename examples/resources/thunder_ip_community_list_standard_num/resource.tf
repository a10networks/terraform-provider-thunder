provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_community_list_standard_num" "thunder_ip_community_list_standard_num" {
  std_list_num = 10
  rules_list {
    std_list_action     = "deny"
    std_list_comm_value = "internet"
  }

}
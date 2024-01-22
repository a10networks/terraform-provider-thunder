provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_large_community_list_standard_num" "thunder_ip_large_community_list_standard_num" {

  std_list_num = 2
  rules_list {
    std_list_lcom_action = "deny"
    std_list_lcomm_value = "12:12:12"
  }

}
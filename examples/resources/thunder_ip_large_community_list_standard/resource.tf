provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_large_community_list_standard" "thunder_ip_large_community_list_standard" {
  standard = "test"
  rules_list {
    standard_lcom_action = "deny"
    standard_lcomm_value = "12:12:12"
  }

}
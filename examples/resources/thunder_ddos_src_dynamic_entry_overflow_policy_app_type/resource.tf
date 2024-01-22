provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_dynamic_entry_overflow_policy_app_type" "thunder_ddos_src_dynamic_entry_overflow_policy_app_type" {
  default_address_type = "ip"
  protocol             = "dns"
  user_tag             = "118"
}
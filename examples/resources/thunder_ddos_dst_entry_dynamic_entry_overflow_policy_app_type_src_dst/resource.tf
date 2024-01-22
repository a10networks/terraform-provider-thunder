provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_dynamic_entry_overflow_policy_app_type_src_dst" "thunder_ddos_dst_entry_dynamic_entry_overflow_policy_app_type_src_dst" {

  protocol       = "dns"
  dummy_name     = "configuration"
  dst_entry_name = "test"
  user_tag       = "74"

}
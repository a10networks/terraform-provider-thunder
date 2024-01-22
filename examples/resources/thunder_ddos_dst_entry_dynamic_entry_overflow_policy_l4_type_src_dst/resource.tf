provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_dynamic_entry_overflow_policy_l4_type_src_dst" "thunder_ddos_dst_entry_dynamic_entry_overflow_policy_l4_type_src_dst" {
  deny           = 1
  glid           = "3"
  protocol       = "tcp"
  dummy_name     = "configuration"
  user_tag       = "test"
  dst_entry_name = "test"
}
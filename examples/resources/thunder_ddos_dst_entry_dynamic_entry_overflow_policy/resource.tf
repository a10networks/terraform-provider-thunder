provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_dynamic_entry_overflow_policy" "thunder_ddos_dst_entry_dynamic_entry_overflow_policy" {

  dst_entry_name = "test"
  bypass         = 1
  dummy_name     = "configuration"
  l4_type_src_dst_list {
    protocol = "tcp"
    deny     = 1
    user_tag = "28"
  }
  log_periodic = 1
  user_tag     = "92"

}
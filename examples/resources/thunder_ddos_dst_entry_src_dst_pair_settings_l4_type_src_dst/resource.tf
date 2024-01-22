provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_dst_pair_settings_l4_type_src_dst" "thunder_ddos_dst_entry_src_dst_pair_settings_l4_type_src_dst" {
  dst_entry_name          = "test"
  max_dynamic_entry_count = 536
  protocol                = "tcp"
  user_tag                = "107"
  all_types               = "all-types"

}
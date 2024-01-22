provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_dst_pair_settings" "thunder_ddos_dst_entry_src_dst_pair_settings" {
  dst_entry_name             = "test"
  age                        = 5
  all_types                  = "all-types"
  enable_class_list_overflow = 1
  l4_type_src_dst_list {
    protocol                = "tcp"
    max_dynamic_entry_count = 22
    user_tag                = "104"
  }
  max_dynamic_entry_count = 26
  user_tag                = "46"

}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy" "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy" {
  dst_entry_name        = "test"
  src_based_policy_name = "test"
  class_list_name       = "test"
  dummy_name            = "configuration"
  bypass                = 1
  exceed_log_cfg {
    log_enable = 1
  }
  log_periodic = 1
  glid         = "3"
  user_tag     = "test"
  l4_type_src_dst_overflow_list {
    protocol = "tcp"
    deny     = 1
    glid     = "3"
    user_tag = "test"
  }

  app_type_src_dst_overflow_list {
    protocol = "dns"
    user_tag = "test"
  }

}
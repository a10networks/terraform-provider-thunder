provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_other_src_based_policy_policy_class_list" "thunder_ddos_dst_zone_port_zone_service_other_src_based_policy_policy_class_list" {
  zone_name               = "test"
  port_other              = "other"
  protocol                = "tcp"
  src_based_policy_name   = "test"
  class_list_name         = "test"
  glid                    = "3"
  action                  = "deny"
  max_dynamic_entry_count = 33
  user_tag                = "test"
  class_list_overflow_policy_list {
    dummy_name = "configuration"
    glid       = "3"
    action     = "deny"
    log_enable = 1
    user_tag   = "test"
  }

}
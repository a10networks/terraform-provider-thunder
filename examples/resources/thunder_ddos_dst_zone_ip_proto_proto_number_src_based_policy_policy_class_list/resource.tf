provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_number_src_based_policy_policy_class_list" "thunder_ddos_dst_zone_ip_proto_proto_number_src_based_policy_policy_class_list" {
  src_based_policy_name   = "test"
  protocol_num            = 80
  zone_name               = "test"
  class_list_name         = "test"
  max_dynamic_entry_count = 20
  user_tag                = "test"
  class_list_overflow_policy_list {
    dummy_name   = "configuration"
    glid         = "3"
    action       = "deny"
    log_enable   = 1
    log_periodic = 1
    user_tag     = "test"
  }


}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_number" "thunder_ddos_dst_zone_ip_proto_proto_number" {
  zone_name                = "testZone"
  age                      = 5
  apply_policy_on_overflow = 1
  deny                     = 1
  drop_frag_pkt            = 1
  dynamic_entry_overflow_policy_list {
    dummy_name = "configuration"
    action     = "bypass"
    user_tag   = "26"
  }
  enable_top_k             = 1
  enable_top_k_destination = 1
  level_list {
    level_num             = "0"
    zone_escalation_score = 193209
    src_escalation_score  = 965485
    user_tag              = "69"
    indicator_list {
      type               = "pkt-rate"
      score              = 473239
      src_threshold_num  = 1039517168
      zone_threshold_num = 773001127
      user_tag           = "61"
    }
  }
  manual_mode_enable = 1
  manual_mode_list {
    config   = "configuration"
    user_tag = "116"
  }
  max_dynamic_entry_count = 33
  protocol_num            = 223
  topk_dst_num_records    = 20
  topk_num_records        = 20

}
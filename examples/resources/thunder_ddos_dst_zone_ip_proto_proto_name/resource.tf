provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_name" "thunder_ddos_dst_zone_ip_proto_proto_name" {
  zone_name                = "testZone"
  age                      = 5
  apply_policy_on_overflow = 1
  deny                     = 1
  drop_frag_pkt            = 1
  dynamic_entry_overflow_policy_list {
    dummy_name = "configuration"
    action     = "bypass"
    user_tag   = "14"
  }

  level_list {
    level_num             = "1"
    zone_escalation_score = 333
    src_escalation_score  = 33
    user_tag              = "test"
    indicator_list {
      type               = "pkt-drop-rate"
      score              = 22
      src_threshold_num  = 222
      zone_threshold_num = 2222
      user_tag           = "test"
    }
  }
  manual_mode_enable = 1
  manual_mode_list {
    config   = "configuration"
    user_tag = "test"
  }
  max_dynamic_entry_count = 444
  protocol                = "icmp-v4"
  topk_dst_num_records    = 20
  topk_num_records        = 20

}
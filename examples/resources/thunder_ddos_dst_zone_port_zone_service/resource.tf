provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service" "thunder_ddos_dst_zone_port_zone_service" {
  zone_name                = "testZone"
  age                      = 5
  apply_policy_on_overflow = 1
  capture_config {
    capture_config_name = "10"
    capture_config_mode = "drop"
  }
  deny = 1
  dynamic_entry_overflow_policy_list {
    dummy_name   = "configuration"
    action       = "bypass"
    log_enable   = 1
    log_periodic = 1
    user_tag     = "8"
  }
  enable_top_k             = 1
  enable_top_k_destination = 1
  level_list {
    level_num                 = "0"
    zone_escalation_score     = 512455
    src_escalation_score      = 562924
    start_pattern_recognition = 1
    user_tag                  = "41"
    indicator_list {
      type               = "pkt-rate"
      score              = 513334
      src_threshold_num  = 1812338350
      zone_threshold_num = 1078849528
      user_tag           = "114"
    }
  }
  manual_mode_enable = 1
  manual_mode_list {
    config   = "configuration"
    user_tag = "36"
  }
  max_dynamic_entry_count = 77
  pattern_recognition {
    algorithm       = "heuristic"
    triggered_by    = "zone-escalation"
    capture_traffic = "all"
  }
  port_num             = 28712
  protocol             = "dns-tcp"
  topk_dst_num_records = 20
  topk_num_records     = 20
}
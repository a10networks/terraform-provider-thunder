provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_range" "thunder_ddos_dst_zone_port_range" {
  zone_name          = "testZone"
  port_range_start   = 22
  port_range_end     = 24
  protocol           = "dns-tcp"
  manual_mode_enable = 1
  deny               = 1
  sflow_tcp {
    sflow_tcp_stateful = 1
  }
  max_dynamic_entry_count    = 432
  apply_policy_on_overflow   = 1
  enable_class_list_overflow = 1
  enable_top_k               = 1
  topk_num_records           = 22
  enable_top_k_destination   = 1
  topk_dst_num_records       = 44
  age                        = 22
  user_tag                   = "test"
  pattern_recognition {
    algorithm          = "heuristic"
    triggered_by       = "packet-rate-exceeds"
    capture_traffic    = "dropped"
    app_payload_offset = 223
  }
  level_list {
    level_num                         = "2"
    zone_escalation_score             = 2222
    src_escalation_score              = 222
    close_sessions_for_unauth_sources = 1
    start_pattern_recognition         = 1
    apply_extracted_filters           = 1
    user_tag                          = "test"
    indicator_list {
      type               = "small-window-ack-rate"
      tcp_window_size    = 22
      score              = 22
      src_threshold_num  = 543
      zone_threshold_num = 986
      user_tag           = "test"
    }
  }
  manual_mode_list {
    config   = "configuration"
    user_tag = "test"
  }
  dynamic_entry_overflow_policy_list {
    dummy_name   = "configuration"
    action       = "deny"
    log_enable   = 1
    log_periodic = 1
    user_tag     = "test"
  }
}
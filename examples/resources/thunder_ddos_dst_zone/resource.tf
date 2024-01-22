provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone" "thunder_ddos_dst_zone" {
  zone_name           = "testZone"
  operational_mode    = "learning"
  continuous_learning = 1
  ip {
    ip_addr = "10.10.10.10"
  }
  ipv6 {
    ip6_addr = "2001:db8:3333:4444:5555:6666:7777:8888"
  }
  description = "test"
  enable_top_k {
    topk_type        = "destination"
    topk_num_records = 23
  }
  drop_frag_pkt         = 1
  advertised_enable     = 1
  inbound_forward_dscp  = 2
  outbound_forward_dscp = 33
  log_enable            = 1
  log_periodic          = 1
  log_high_frequency    = 1
  rate_limit            = 2
  user_tag              = "test"
  detection {
    settings = "settings"
    notification {
      configuration = "configuration"
    }
    outbound_detection {
      configuration = "configuration"
      toggle        = "enable"
      indicator_list {
        type          = "pkt-rate"
        threshold_num = 22
        user_tag      = "test"
      }
    }
    service_discovery {
      configuration      = "configuration"
      toggle             = "disable"
      pkt_rate_threshold = 22
    }
    packet_anomaly_detection {
      configuration = "configuration"
      toggle        = "enable"
      indicator_list {
        type          = "port-zero-pkt-rate"
        threshold_num = 22
        user_tag      = "test"
      }
    }
    victim_ip_detection {
      configuration    = "configuration"
      toggle           = "enable"
      histogram_toggle = "histogram-enable"
      indicator_list {
        type             = "pkt-rate"
        ip_threshold_num = 22
        user_tag         = "test"
      }
    }
  }
  capture_config_list {
    name = "10"
    mode = "all"
  }
  ip_proto {
    proto_tcp_udp_list {
      protocol      = "udp"
      deny          = 1
      drop_frag_pkt = 1
    }
  }
  port {
    zone_service_other_list {
      port_other                 = "other"
      protocol                   = "tcp"
      manual_mode_enable         = 1
      enable_top_k               = 1
      topk_num_records           = 22
      enable_top_k_destination   = 1
      topk_dst_num_records       = 33
      sflow_packets              = 1
      max_dynamic_entry_count    = 333
      apply_policy_on_overflow   = 1
      enable_class_list_overflow = 1
      age                        = 3
      pattern_recognition {
        algorithm       = "heuristic"
        triggered_by    = "packet-rate-exceeds"
        capture_traffic = "all"
      }
      level_list {
        level_num                         = "1"
        zone_escalation_score             = 333
        src_escalation_score              = 33
        close_sessions_for_unauth_sources = 1
        start_pattern_recognition         = 1
        apply_extracted_filters           = 1
        user_tag                          = "test"
        indicator_list {
          type               = "pkt-drop-rate"
          score              = 22
          src_threshold_num  = 222
          zone_threshold_num = 2222
          user_tag           = "test"
        }
      }
      manual_mode_list {
        config   = "configuration"
        user_tag = "test"
      }
    }
  }
  port_range_list {
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

}
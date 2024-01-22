provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_other" "thunder_ddos_dst_zone_port_zone_service_other" {
  zone_name                = "test"
  port_other               = "other"
  protocol                 = "tcp"
  manual_mode_enable       = 1
  enable_top_k             = 1
  topk_num_records         = 22
  enable_top_k_destination = 1
  topk_dst_num_records     = 33
  deny                     = 1
  glid_cfg {
    glid = "3"
  }
  sflow_tcp {
    sflow_tcp_basic = 1
  }
  max_dynamic_entry_count    = 20
  enable_class_list_overflow = 1
  age                        = 20
  pattern_recognition {
    algorithm       = "heuristic"
    triggered_by    = "zone-escalation"
    capture_traffic = "all"
  }
  level_list {
    level_num                         = "1"
    src_default_glid                  = "3"
    zone_escalation_score             = 3
    src_escalation_score              = 44
    close_sessions_for_unauth_sources = 1
    start_pattern_recognition         = 1
    apply_extracted_filters           = 1
    user_tag                          = "testing"
    indicator_list {
      type               = "pkt-rate"
      score              = 2
      src_threshold_num  = 33
      zone_threshold_num = 22
      user_tag           = "test"
    }

  }

  manual_mode_list {
    config           = "configuration"
    src_default_glid = "3"
    user_tag         = "testing"
  }

  src_based_policy_list {
    src_based_policy_name = "test"
    user_tag              = "testing"
    policy_class_list_list {
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

  }

  dynamic_entry_overflow_policy_list {
    dummy_name   = "configuration"
    glid         = "3"
    action       = "deny"
    log_enable   = 1
    log_periodic = 1
    user_tag     = "testing"
  }


}
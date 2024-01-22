provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry" "thunder_ddos_dst_entry" {

  dst_entry_name = "test"
  ip_addr        = "10.10.10.11"
  description    = "testtt"
  exceed_log_cfg {
    log_enable         = 1
    log_with_sflow     = 1
    log_high_frequency = 1
  }
  log_periodic               = 1
  drop_frag_pkt              = 1
  drop_on_no_src_dst_default = 1
  blackhole_on_glid_exceed   = 2
  dest_nat_ip                = "10.11.11.11"
  drop_disable               = 1
  drop_disable_fwd_immediate = 1
  operational_mode           = "bypass"
  advertised_enable          = 1
  inbound_forward_dscp       = 3
  enable_top_k {
    topk_type        = "destination"
    topk_num_records = 22
  }

  user_tag = "maintest"
  capture_config_list {
    name = "10"
    mode = "drop"
  }

  l4_type_list {
    protocol                              = "tcp"
    deny                                  = 1
    max_rexmit_syn_per_flow               = 2
    max_rexmit_syn_per_flow_exceed_action = "drop"
    syn_auth                              = "send-rst"
    tcp_reset_client                      = 1
    tcp_reset_server                      = 1
    drop_on_no_port_match                 = "enable"
    drop_frag_pkt                         = 1
    undefined_port_hit_statistics {
      undefined_port_hit_statistics = 1
    }
    user_tag = "pqrs"
  }

  port_list {
    port_num         = 22
    protocol         = "tcp"
    detection_enable = 1
    enable_top_k     = 1
    topk_num_records = 22
    deny             = 1
    sflow {
      polling {
        sflow_tcp {
          sflow_tcp_basic = 1
        }
      }
    }

    user_tag = "abcd"
  }

  src_port_list {
    port_num = 22
    protocol = "tcp"
    deny     = 1

    user_tag = "test"
  }

  ip_proto_list {
    port_num = 22
    deny     = 1

    user_tag = "testtt"
  }

  src_dst_pair {
    default = 1
    bypass  = 1
    exceed_log_cfg {
      log_enable = 1
    }
    log_periodic = 1

    l4_type_src_dst_list {
      protocol = "tcp"
      deny     = 1

      user_tag = "testuser"
    }

    app_type_src_dst_list {
      protocol = "dns"

      user_tag = "testuser"
    }

  }
  src_dst_pair_settings_list {
    all_types                  = "all-types"
    age                        = 22
    max_dynamic_entry_count    = 22
    enable_class_list_overflow = 1

    user_tag = "test"
    l4_type_src_dst_list {
      protocol                = "tcp"
      max_dynamic_entry_count = 22

      user_tag = "test"
    }

  }

  dynamic_entry_overflow_policy_list {
    dummy_name   = "configuration"
    bypass       = 1
    log_periodic = 1

    user_tag = "testing"
    l4_type_src_dst_list {
      protocol = "tcp"
      deny     = 1

      user_tag = "test"
    }

    app_type_src_dst_list {
      protocol = "dns"

      user_tag = "test"
    }

  }

}
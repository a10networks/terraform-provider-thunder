provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_tcp" "thunder_ddos_template_tcp" {
  name = "test"
  action_cfg {
    action_on_ack      = 1
    reset              = 1
    timeout            = 3
    min_retry_gap      = 2
    authenticate_only  = 1
    rto_authentication = 1
  }
  action_on_syn_rto_retry_count = 3
  action_on_ack_rto_retry_count = 3
  age                           = 4
  black_list_out_of_seq         = 33
  black_list_retransmit         = 33
  black_list_zero_win           = 44
  syn_auth                      = "send-rst"
  conn_rate_limit_on_syn_only   = 1
  per_conn_rate_interval        = "1sec"
  per_conn_pkt_rate_limit       = 3
  per_conn_pkt_rate_action      = "drop"
  dst {
    rate_limit {
      syn_rate_limit {
        dst_syn_rate_action = "drop"
      }
    }
  }
  src {
    rate_limit {
      syn_rate_limit {
        src_syn_rate_action = "drop"
      }
    }
  }
  allow_synack_skip_authentications = 1
  synack_rate_limit                 = 3
  allow_syn_otherflags              = 1
  allow_tcp_tfo                     = 1
  drop_known_resp_src_port_cfg {
    drop_known_resp_src_port = 1
    exclude_src_resp_port    = 1
  }
  tunnel_encap {
    gre_cfg {
      gre_encap = 1
      gre_always {
        gre_ipv4              = "10.10.10.10"
        preserve_src_ipv4_gre = 1
      }
    }
  }

  user_tag = "test"
  progression_tracking {
    progression_tracking_enabled = "enable-check"
    violation                    = 3
    progression_tracking_action  = "drop"

    connection_tracking {
      progression_tracking_conn_enabled = "enable-check"
      conn_sent_max                     = 10
      conn_sent_min                     = 8
      conn_rcvd_max                     = 10
      conn_rcvd_min                     = 8
      conn_rcvd_sent_ratio_min          = 18
      conn_rcvd_sent_ratio_max          = 20
      conn_duration_max                 = 22
      conn_duration_min                 = 3
      conn_violation                    = 3
      progression_tracking_conn_action  = "drop"

    }
  }
  filter_list {
    tcp_filter_seq       = 3
    tcp_filter_regex     = "test"
    tcp_filter_unmatched = 1
    tcp_filter_action    = "count-only"

    user_tag = "test"
  }

}
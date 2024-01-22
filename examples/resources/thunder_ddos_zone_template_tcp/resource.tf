provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_tcp" "thunder_ddos_zone_template_tcp" {
  name = "test"
  age  = 33
  out_of_seq_cfg {
    out_of_seq        = 3
    out_of_seq_action = "ignore"
  }
  max_rexmit_syn_per_flow_cfg {
    max_rexmit_syn_per_flow        = 3
    max_rexmit_syn_per_flow_action = "drop"
  }
  per_conn_retransmit_rate_cfg {
    per_conn_retransmit_rate_limit  = 3
    per_conn_retransmit_rate_action = "drop"
  }
  per_conn_zero_win_rate_cfg {
    per_conn_zero_win_rate_limit  = 3
    per_conn_zero_win_rate_action = "ignore"
  }
  per_conn_pkt_rate_cfg {
    per_conn_pkt_rate_limit  = 3
    per_conn_pkt_rate_action = "drop"
  }
  per_conn_rate_interval = "1sec"
  dst {
    rate_limit {
      syn_rate_limit {
        dst_syn_rate_action = "ignore"
      }
    }
  }
  src {
    rate_limit {
      syn_rate_limit {
        src_syn_rate_limit  = 3
        src_syn_rate_action = "drop"
      }
    }
  }
  allow_synack_skip_authentications = 1
  synack_rate_limit                 = 33
  allow_syn_otherflags              = 1
  allow_tcp_tfo                     = 1
  conn_rate_limit_on_syn_only       = 1
  action_on_syn_rto_retry_count     = 3
  action_on_ack_rto_retry_count     = 3
  known_resp_src_port_cfg {
    known_resp_src_port        = 1
    known_resp_src_port_action = "ignore"
  }
  syn_authentication {
    syn_auth_timeout     = 22
    syn_auth_min_delay   = 2
    syn_auth_pass_action = "authenticate-src"
    syn_auth_fail_action = "drop"
  }
  ack_authentication {
    ack_auth_timeout     = 3
    ack_auth_min_delay   = 3
    ack_auth_only        = 1
    ack_auth_rto         = 1
    ack_auth_pass_action = "authenticate-src"
    ack_auth_fail_action = "drop"
  }
  user_tag = "test"
  progression_tracking {
    progression_tracking_enabled     = "enable-check"
    request_response_model           = "enable"
    violation                        = 33
    ignore_tls_handshake             = 1
    response_length_max              = 20
    request_length_min               = 20
    request_length_max               = 40
    response_request_min_ratio       = 20
    response_request_max_ratio       = 40
    first_request_max_time           = 3
    request_to_response_max_time     = 30
    response_to_request_max_time     = 3
    profiling_request_response_model = 1
    profiling_connection_life_model  = 1
    progression_tracking_action      = "drop"
    connection_tracking {
      progression_tracking_conn_enabled = "enable-check"
      conn_sent_max                     = 20
      conn_sent_min                     = 10
      conn_rcvd_max                     = 20
      conn_rcvd_min                     = 10
      conn_rcvd_sent_ratio_min          = 15
      conn_rcvd_sent_ratio_max          = 30
      conn_duration_max                 = 44
      conn_duration_min                 = 22
      conn_violation                    = 5
      progression_tracking_conn_action  = "drop"
    }
    time_window_tracking {
      progression_tracking_win_enabled    = "enable-check"
      window_sent_max                     = 20
      window_sent_min                     = 10
      window_rcvd_max                     = 20
      window_rcvd_min                     = 10
      window_rcvd_sent_ratio_min          = 10
      window_rcvd_sent_ratio_max          = 20
      window_violation                    = 2
      progression_tracking_windows_action = "drop"
    }
  }
  filter_list {
    tcp_filter_name          = "test"
    tcp_filter_seq           = 3
    tcp_filter_regex         = "test"
    tcp_filter_inverse_match = 1
    tcp_filter_action        = "drop"

    user_tag = "test"
  }
}
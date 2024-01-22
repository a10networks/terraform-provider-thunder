provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_udp" "thunder_ddos_zone_template_udp" {
  name = "test"
  age  = 10
  per_conn_pkt_rate_cfg {
    per_conn_pkt_rate_limit  = 33
    per_conn_pkt_rate_action = "ignore"
  }
  per_conn_rate_interval                = "1sec"
  spoof_detect_min_delay_interval       = "1sec"
  spoof_detect_min_delay                = 3
  spoof_detect_pass_action              = "authenticate-src"
  spoof_detect_fail_action              = "drop"
  token_authentication                  = 1
  token_authentication_salt_prefix      = 1
  token_authentication_salt_prefix_curr = 33
  token_authentication_salt_prefix_prev = 22
  token_authentication_formula          = "sha1_Salt-UintDstIp-DstPort"
  previous_salt_timeout                 = 2
  token_authentication_public_address   = 1
  public_ipv4_addr                      = "10.10.10.10"
  known_resp_src_port_cfg {
    known_resp_src_port        = 1
    known_resp_src_port_action = "ignore"
    exclude_src_resp_port      = 1
  }
  ntp_monlist_cfg {
    ntp_monlist        = 1
    ntp_monlist_action = "ignore"
  }
  max_payload_size_cfg {
    max_payload_size = 40
  }
  min_payload_size_cfg {
    min_payload_size = 20
  }
  user_tag = "test"
  filter_list {
    udp_filter_name          = "test"
    udp_filter_seq           = 3
    udp_filter_regex         = "test"
    udp_filter_inverse_match = 1
    udp_filter_action        = "drop"
    user_tag                 = "test"
  }
}
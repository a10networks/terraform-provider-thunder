provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_udp" "thunder_ddos_template_udp" {
  name = "test"
  age  = 8
  drop_known_resp_src_port_cfg {
    drop_known_resp_src_port = 1
    exclude_src_resp_port    = 1
  }
  drop_ntp_monlist = 1
  filter_list {
    udp_filter_seq       = 1
    udp_filter_regex     = "1209"
    udp_filter_unmatched = 1
    udp_filter_action    = "blacklist-src"
    user_tag             = "16"
  }
  max_payload_size        = 1413
  min_payload_size        = 1054
  per_conn_pkt_rate_limit = 1246
  per_conn_rate_interval  = "1sec"
  previous_salt_timeout   = 1
  spoof_detect_cfg {
    spoof_detect                        = 1
    min_retry_gap_interval              = "1sec"
    spoof_detect_retry_timeout_val_only = 5
    min_retry_gap                       = 3
  }
  tunnel_encap {
    ip_encap = 1
    always {
      ipv4_addr         = "10.10.10.10"
      preserve_src_ipv4 = 1
    }
  }
  user_tag = "31"

}
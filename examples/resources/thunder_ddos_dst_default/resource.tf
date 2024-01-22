provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_default" "thunder_ddos_dst_default" {
  age                        = 6
  apply_policy_on_overflow   = 1
  default_address_type       = "ip"
  deny                       = 1
  disable                    = 1
  drop_disable               = 1
  drop_disable_fwd_immediate = 1
  drop_frag_pkt              = 1
  exceed_log_cfg {
    log_enable        = 1
    with_sflow_sample = 1
  }
  inbound_forward_dscp = 22
  ip_proto_list {
    port_num = 111
    deny     = 1
    user_tag = "107"
  }
  l4_type_list {
    protocol                = "tcp"
    deny                    = 1
    max_rexmit_syn_per_flow = 6
    syn_auth                = "send-rst"
    tcp_reset_client        = 1
    tcp_reset_server        = 1
    drop_on_no_port_match   = "enable"
    drop_frag_pkt           = 1
    user_tag                = "32"
  }
  log_periodic            = 1
  max_dynamic_entry_count = 20
  port_list {
    port_num = 700
    protocol = "tcp"
    deny     = 1
    user_tag = "108"
  }
  src_port_list {
    port_num = 2267
    protocol = "tcp"
    deny     = 1
    user_tag = "61"
  }
  user_tag = "115"
}
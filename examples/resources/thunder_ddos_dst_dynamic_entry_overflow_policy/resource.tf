provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_dynamic_entry_overflow_policy" "thunder_ddos_dst_dynamic_entry_overflow_policy" {
  default_address_type = "ip"
  exceed_log_cfg {
    log_enable        = 1
    with_sflow_sample = 1
  }
  drop_disable               = 1
  drop_disable_fwd_immediate = 1
  log_periodic               = 1
  inbound_forward_dscp       = 3
  outbound_forward_dscp      = 4
  glid                       = "3"
  user_tag                   = "test"
  l4_type_list {
    protocol                = "tcp"
    glid                    = "3"
    deny                    = 1
    max_rexmit_syn_per_flow = 3
    syn_auth                = "send-rst"
    tcp_reset_client        = 1
    tcp_reset_server        = 1
    drop_on_no_port_match   = "enable"
    drop_frag_pkt           = 1
    user_tag                = "test"
  }

  port_list {
    port_num = 80
    protocol = "tcp"
    deny     = 1
    glid     = "3"
    user_tag = "test"
  }

  src_port_list {
    port_num = 44
    protocol = "tcp"
    deny     = 1
    glid     = "3"
    user_tag = "test"
  }

  ip_proto_list {
    port_num = 32
    deny     = 1
    glid     = "3"
    user_tag = "test"
  }
}
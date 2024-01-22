provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_l4_type" "thunder_ddos_dst_entry_l4_type" {
  dst_entry_name                        = "test"
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
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_default_l4_type" "thunder_ddos_dst_default_l4_type" {
  default_address_type  = "ip"
  deny                  = 1
  drop_frag_pkt         = 1
  drop_on_no_port_match = "enable"
  glid                  = "3"
  protocol              = "tcp"
  syn_auth              = "send-rst"
  tcp_reset_client      = 1
  tcp_reset_server      = 1
  user_tag              = "test"

}
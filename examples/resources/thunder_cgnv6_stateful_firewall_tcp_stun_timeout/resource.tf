provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_tcp_stun_timeout" "thunder_cgnv6_stateful_firewall_tcp_stun_timeout" {

  port                        = 22
  port_end                    = 26
  stun_timeout_val_port_range = 4

}
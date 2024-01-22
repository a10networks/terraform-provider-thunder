provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_udp_stun_timeout" "thunder_cgnv6_stateful_firewall_udp_stun_timeout" {
  port                        = 24
  port_end                    = 26
  stun_timeout_val_port_range = 3
}
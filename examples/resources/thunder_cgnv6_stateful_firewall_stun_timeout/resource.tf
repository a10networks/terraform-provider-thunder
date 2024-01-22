provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_stun_timeout" "thunder_cgnv6_stateful_firewall_stun_timeout" {
  port                        = 20
  port_end                    = 24
  stun_timeout_val_port_range = 4
}
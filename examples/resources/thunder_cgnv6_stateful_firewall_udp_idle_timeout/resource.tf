provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_udp_idle_timeout" "thunder_cgnv6_stateful_firewall_udp_idle_timeout" {

  idle_timeout_val_port_range = 302
  port                        = 20
  port_end                    = 22

}
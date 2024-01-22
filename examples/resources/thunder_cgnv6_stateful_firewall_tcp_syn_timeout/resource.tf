provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_tcp_syn_timeout" "thunder_cgnv6_stateful_firewall_tcp_syn_timeout" {

  syn_timeout_val = 5

}
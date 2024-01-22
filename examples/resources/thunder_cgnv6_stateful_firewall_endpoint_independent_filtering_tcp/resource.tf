provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_endpoint_independent_filtering_tcp" "thunder_cgnv6_stateful_firewall_endpoint_independent_filtering_tcp" {
  port_list {
    port     = 123
    port_end = 1234
  }
}
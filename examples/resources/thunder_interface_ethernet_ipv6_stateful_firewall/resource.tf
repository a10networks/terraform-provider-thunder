provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_ipv6_stateful_firewall" "thunder_interface_ethernet_ipv6_stateful_firewall" {
  ifnum       = 2
  access_list = 0
  inside      = 1
  outside     = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_ipv6_stateful_firewall" "thunder_interface_trunk_ipv6_stateful_firewall" {

  ifnum       = 11
  access_list = 0
  inside      = 1
  outside     = 1
}
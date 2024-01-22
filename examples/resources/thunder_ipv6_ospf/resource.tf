provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_ospf" "thunder_ipv6_ospf" {
  display_route_single_line = 0
}
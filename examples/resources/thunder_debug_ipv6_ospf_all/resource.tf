provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ipv6_ospf_all" "thunder_debug_ipv6_ospf_all" {
  dumy = 0
}
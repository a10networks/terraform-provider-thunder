provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ipv6_ospf_nsm" "thunder_debug_ipv6_ospf_nsm" {
  interface    = 0
  redistribute = 0
}
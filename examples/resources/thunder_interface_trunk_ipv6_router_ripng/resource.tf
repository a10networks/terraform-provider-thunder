provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_ipv6_router_ripng" "thunder_interface_trunk_ipv6_router_ripng" {

  ifnum = 11
  rip   = 1
}
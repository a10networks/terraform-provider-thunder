provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_loopback_ipv6_router_ripng" "thunder_interface_loopback_ipv6_router_ripng" {

  ifnum = 4
  rip   = 1
}
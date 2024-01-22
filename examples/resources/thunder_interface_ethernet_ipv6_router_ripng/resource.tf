provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_ipv6_router_ripng" "thunder_interface_ethernet_ipv6_router_ripng" {
  ifnum = 2
  rip   = 1
}
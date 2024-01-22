provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif_ipv6_router_ripng" "thunder_interface_lif_ipv6_router_ripng" {

  ifname = "test"
  rip    = 1
}
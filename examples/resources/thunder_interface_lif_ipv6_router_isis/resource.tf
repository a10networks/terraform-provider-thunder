provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif_ipv6_router_isis" "thunder_interface_lif_ipv6_router_isis" {

  ifname = "test"
  tag    = "18"
}
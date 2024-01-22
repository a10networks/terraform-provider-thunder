provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_address_family_ipv6_network_monitor_default" "thunder_router_bgp_address_family_ipv6_network_monitor_default" {

  as_number               = "300"
  network_monitor_default = 0
}
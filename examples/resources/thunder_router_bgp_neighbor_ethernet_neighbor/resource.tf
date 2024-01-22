provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_neighbor_ethernet_neighbor" "thunder_router_bgp_neighbor_ethernet_neighbor" {

  as_number       = "300"
  ethernet        = 2
  unnumbered      = 1
  peer_group_name = "A10"
}
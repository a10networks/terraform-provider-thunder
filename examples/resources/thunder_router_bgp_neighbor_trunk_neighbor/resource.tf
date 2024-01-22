provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_neighbor_peer_group_neighbor" "thunder_router_bgp_neighbor_peer_group_neighbor" {

  peer_group     = "A10"
  peer_group_key = 1
  activate       = 1
  as_number      = "300"
}
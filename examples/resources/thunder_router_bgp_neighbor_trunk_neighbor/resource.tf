provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_bgp" "bgp" {
  as_number = "300"
}
resource "thunder_router_bgp_neighbor_peer_group_neighbor" "PeerGroup" {
  peer_group     = "A10"
  peer_group_key = 1
  activate       = 1
  as_number      = "300"
}

resource "thunder_router_bgp_neighbor_trunk_neighbor" "TrunkBGP" {
  as_number       = "300"
  process_id      = 200
  trunk           = 5
  unnumbered      = 1
  peer_group_name = "A10"
}
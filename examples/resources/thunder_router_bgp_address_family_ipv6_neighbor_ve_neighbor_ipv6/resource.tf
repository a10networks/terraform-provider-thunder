provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_address_family_ipv6_neighbor_trunk_neighbor_ipv6" "thunderRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Test" {

  trunk           = 5
  peer_group_name = "A10"
  as_number       = "300"
  sequence        = 1
  process_id      = 200
}
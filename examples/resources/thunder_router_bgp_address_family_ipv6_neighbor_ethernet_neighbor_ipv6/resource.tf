provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_bgp_address_family_ipv6_neighbor_ethernet_neighbor_ipv6" "BgpAddressFamilyIpv6NeigEthNeiIpv6Test" {
  ethernet        = 3
  peer_group_name = "A12"
  as_number       = "300"
  process_id      = 200
}
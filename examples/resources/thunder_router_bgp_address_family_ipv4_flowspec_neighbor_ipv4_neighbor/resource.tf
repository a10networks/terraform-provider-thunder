provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_address_family_ipv4_flowspec_neighbor_ipv4_neighbor" "thunder_router_bgp_address_family_ipv4_flowspec_neighbor_ipv4_neighbor" {

  as_number     = "300"
  activate      = 0
  neighbor_ipv4 = "10.10.10.10"
  neighbor_route_map_lists {
    nbr_route_map      = "17"
    nbr_rmap_direction = "in"
  }
  send_community_val = "both"
}
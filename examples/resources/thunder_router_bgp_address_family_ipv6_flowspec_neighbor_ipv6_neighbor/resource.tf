provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_address_family_ipv6_flowspec_neighbor_ipv6_neighbor" "thunder_router_bgp_address_family_ipv6_flowspec_neighbor_ipv6_neighbor" {

  as_number     = "300"
  activate      = 0
  neighbor_ipv6 = "2001:db8:3333:4444:5555:6666:7777:8888"
  neighbor_route_map_lists {
    nbr_route_map      = "25"
    nbr_rmap_direction = "in"
  }
  send_community_val = "both"
}
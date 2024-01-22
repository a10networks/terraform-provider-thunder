provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_address_family_ipv6_flowspec" "thunder_router_bgp_address_family_ipv6_flowspec" {

  as_number = "300"
  neighbor {
    ipv4_neighbor_list {
      neighbor_ipv4 = "10.10.10.10"
      activate      = 0
      neighbor_route_map_lists {
        nbr_route_map      = "87"
        nbr_rmap_direction = "in"
      }
      send_community_val = "both"
    }
    ipv6_neighbor_list {
      neighbor_ipv6 = "2001:db8:3333:4444:5555:6666:7777:8888"
      activate      = 0
      neighbor_route_map_lists {
        nbr_route_map      = "52"
        nbr_rmap_direction = "in"
      }
      send_community_val = "both"
    }
  }
}
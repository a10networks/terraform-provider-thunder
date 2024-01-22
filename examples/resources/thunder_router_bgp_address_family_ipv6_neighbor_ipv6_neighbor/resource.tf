provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_address_family_ipv6_neighbor_ipv6_neighbor" "AFIpv6NeighborIpv6NeighborTest" {
  as_number         = "300"
  neighbor_ipv6     = "2222::1"
  activate          = 1
  allowas_in        = 1
  allowas_in_count  = 10
  graceful_restart  = 1
  default_originate = 1
  route_map         = "Ipv6RouteMap"
  distribute_lists {
    distribute_list           = "DistIpv6BGP"
    distribute_list_direction = "in"
  }

  neighbor_filter_lists {
    filter_list           = "FilterBgp"
    filter_list_direction = "in"
  }
  maximum_prefix       = 100
  maximum_prefix_thres = 80
  next_hop_self        = 1

  remove_private_as = 1

  neighbor_route_map_lists {
    nbr_route_map      = "Ipv6RouteMap"
    nbr_rmap_direction = "in"
  }
  send_community_val = "standard"
  inbound            = 1
  unsuppress_map     = "UnsuppIpv6"
  weight             = 300
}
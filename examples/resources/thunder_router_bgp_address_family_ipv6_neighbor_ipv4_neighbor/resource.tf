provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_address_family_ipv6_neighbor_ipv4_neighbor" "ipv6neighIpv4neigh" {
  as_number         = "300"
  neighbor_ipv4     = "1.1.1.1"
  activate          = 1
  allowas_in        = 1
  allowas_in_count  = 4
  graceful_restart  = 0
  default_originate = 1
  route_map         = "BGPIn"
  distribute_lists {
    distribute_list           = "BGPDistribute"
    distribute_list_direction = "in"
  }
  neighbor_filter_lists {
    filter_list           = "BGPFilter"
    filter_list_direction = "in"
  }
  maximum_prefix       = 300
  maximum_prefix_thres = 80
  next_hop_self        = 1
  remove_private_as    = 0
  unsuppress_map       = "unsuppBGP"
  neighbor_route_map_lists {
    nbr_route_map      = "bgprouteap"
    nbr_rmap_direction = "in"
  }
  send_community_val = "standard"
  inbound            = 0
  weight             = 300
}
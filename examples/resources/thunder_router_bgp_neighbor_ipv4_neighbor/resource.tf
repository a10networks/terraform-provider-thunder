provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}


resource "thunder_router_bgp_neighbor_ipv4_neighbor" "resourceRouterBgpNeighborIpv4NeighborTest" {
  as_number = "300"
  distribute_lists {
    distribute_list           = "BGPDistribute"
    distribute_list_direction = "in"
  }
  neighbor_ipv4              = "99.99.99.99"
  nbr_remote_as              = 300
  strict_capability_match    = 0
  activate                   = 1
  advertisement_interval     = 500
  allowas_in                 = 1
  allowas_in_count           = 4
  as_origination_interval    = 500
  dynamic                    = 1
  prefix_list_direction      = "both"
  route_refresh              = 1
  graceful_restart           = 1
  collide_established        = 1
  default_originate          = 1
  route_map                  = "BGPIn"
  description                = "BGPRouter"
  disallow_infinite_holdtime = 1
  acos_application_only      = 1
  dont_capability_negotiate  = 1
  ebgp_multihop              = 1
  ebgp_multihop_hop_count    = 10
  enforce_multihop           = 1
  bfd                        = 1
  multihop                   = 1
  key_id                     = 40
  key_type                   = "md5"
  bfd_value                  = "string"
  neighbor_filter_lists {
    filter_list           = "BGPFilter"
    filter_list_direction = "in"
  }
  maximum_prefix       = 300
  maximum_prefix_thres = 80
  next_hop_self        = 1
  override_capability  = 1
  pass_value           = "BGPPASS"
  passive              = 1
  remove_private_as    = 1
  neighbor_route_map_lists {
    nbr_route_map      = "bgprouteap"
    nbr_rmap_direction = "in"
  }
  send_community_val = "standard"
  inbound            = 1
  shutdown           = 1

  timers_keepalive = 4000
  timers_holdtime  = 500
  connect          = 1
  unsuppress_map   = "unsuppBGP"
  update_source_ip = "11.11.11.11"
  weight           = 300
}
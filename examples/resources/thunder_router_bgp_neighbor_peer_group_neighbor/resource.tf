provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_bgp_neighbor_peer_group_neighbor" "BgpNeighborPeerGroupNei" {
  as_number                  = "300"
  peer_group                 = "A10"
  peer_group_key             = 1
  peer_group_remote_as       = 400
  activate                   = 1
  advertisement_interval     = 10
  allowas_in                 = 1
  allowas_in_count           = 10
  as_origination_interval    = 20
  dynamic                    = 1
  route_refresh              = 1
  extended_nexthop           = 0
  collide_established        = 1
  default_originate          = 1
  route_map                  = "Ipv6RouteMap"
  description                = "PEERGroup"
  disallow_infinite_holdtime = 1
  distribute_lists {
    distribute_list           = "DistIpv6BGP"
    distribute_list_direction = "in"
  }
  dont_capability_negotiate = 1
  ebgp_multihop             = 1
  ebgp_multihop_hop_count   = 100
  enforce_multihop          = 1
  bfd                       = 1
  multihop                  = 1
  neighbor_filter_lists {
    filter_list           = "FilterBgp"
    filter_list_direction = "in"
  }
  maximum_prefix       = 100
  maximum_prefix_thres = 80
  next_hop_self        = 1
  override_capability  = 0
  pass_value           = "PassPeergroup"
  passive              = 1

  remove_private_as = 1
  neighbor_route_map_lists {
    nbr_route_map      = "Ipv6RouteMap"
    nbr_rmap_direction = "in"
  }
  send_community_val      = "standard"
  inbound                 = 1
  shutdown                = 1
  strict_capability_match = 0
  timers_holdtime         = 200
  timers_keepalive        = 200
  connect                 = 1
  unsuppress_map          = "UnsuppIpv6"
  update_source_ip        = "1.1.1.1"
  weight                  = 300
}

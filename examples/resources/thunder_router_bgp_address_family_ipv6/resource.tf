provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_bgp_address_family_ipv6" "resourceRouterBgpAddressFamilyIpv6Test" {
  as_number = "300"
  bgp {
    dampening                = 1
    dampening_half           = 1
    dampening_start_reuse    = 1
    dampening_start_supress  = 1
    dampening_max_supress    = 1
    dampening_unreachability = 1
  }
  distance {
    distance_ext   = 1
    distance_int   = 1
    distance_local = 1
  }
  maximum_paths_value = 2
  auto_summary        = 1
  synchronization     = 1
  originate           = 1
  network {
    synchronization {
      network_synchronization = 1
    }
    ipv6_network_list {
      network_ipv6 = "1111::1/64"
      route_map    = "ipv6Routemap"
      backdoor     = 1
      description  = "AddIpv6"
      comm_value   = "internet"
    }
  }
  aggregate_address_list {
    aggregate_address = "3333::1/64"
    as_set            = 1
    summary_only      = 1
  }

  redistribute {
    connected_cfg {
      connected = 1
      route_map = "ipv6"
    }
    floating_ip_cfg {
      floating_ip = 1
      route_map   = "ipv6"
    }
    nat64_cfg {
      nat64     = 1
      route_map = "ipv6"
    }
    nat_map_cfg {
      nat_map   = 1
      route_map = "ipv6"
    }
    lw4o6_cfg {
      lw4o6     = 1
      route_map = "ipv6"
    }
    static_nat_cfg {
      static_nat = 1
      route_map  = "ipv6"
    }
    ip_nat_cfg {
      ip_nat    = 1
      route_map = "ipv6"
    }
    ip_nat_list_cfg {
      ip_nat_list = 1
      route_map   = "ipv6"
    }
    isis_cfg {
      isis      = 1
      route_map = "ipv6"
    }
    ospf_cfg {
      ospf      = 1
      route_map = "ipv6"
    }
    rip_cfg {
      rip       = 1
      route_map = "ipv6"
    }
    static_cfg {
      static    = 1
      route_map = "ipv6"
    }
    vip {
      only_flagged_cfg {
        only_flagged = 1
        route_map    = "ipv6"
      }
      only_not_flagged_cfg {
        only_not_flagged = 1
        route_map        = "ipv6"
      }
    }

  }
  neighbor {
    ipv4_neighbor_list {
      neighbor_ipv4    = "1.1.1.1"
      peer_group_name  = "A10"
      activate         = 1
      allowas_in       = 1
      allowas_in_count = 10
      neighbor_filter_lists {
        filter_list           = "FilterBgp"
        filter_list_direction = "in"
      }
      maximum_prefix       = 100
      maximum_prefix_thres = 80
      neighbor_prefix_lists {
        nbr_prefix_list           = "prefixIpv6"
        nbr_prefix_list_direction = "in"
      }

      neighbor_route_map_lists {
        nbr_route_map      = "Ipv6RouteMap"
        nbr_rmap_direction = "in"
      }
      inbound = 1
      weight  = 300
    }
    ipv6_neighbor_list {
      neighbor_ipv6    = "2222::1"
      peer_group_name  = "A10"
      activate         = 1
      allowas_in       = 1
      allowas_in_count = 10
      neighbor_filter_lists {
        filter_list           = "FilterBgp"
        filter_list_direction = "in"
      }
      maximum_prefix       = 100
      maximum_prefix_thres = 80
      neighbor_prefix_lists {
        nbr_prefix_list           = "prefixIpv6"
        nbr_prefix_list_direction = "in"
      }

      neighbor_route_map_lists {
        nbr_route_map      = "Ipv6RouteMap"
        nbr_rmap_direction = "in"
      }
      inbound = 1
      weight  = 300
    }
    ethernet_neighbor_ipv6_list {
      ethernet        = 4
      peer_group_name = "A10"
    }
    ve_neighbor_ipv6_list {
      ve              = 300
      peer_group_name = "A10"
    }
    peer_group_neighbor_list {
      peer_group       = "A10"
      activate         = 1
      allowas_in       = 1
      allowas_in_count = 10

      maximum_prefix       = 100
      maximum_prefix_thres = 80
      next_hop_self        = 1
      remove_private_as    = 1
      neighbor_route_map_lists {
        nbr_route_map      = "Ipv6RouteMap"
        nbr_rmap_direction = "in"
      }
      inbound = 1

      weight = 300
    }
  }
}


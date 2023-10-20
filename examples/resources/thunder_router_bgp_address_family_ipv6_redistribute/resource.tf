provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}


resource "thunder_router_bgp_address_family_ipv6_redistribute" "AFIpv6RedistributeTest" {

  as_number = "300"
  connected_cfg {
    connected = 1
    route_map = "ipv6"
  }
  floating_ip_cfg {
    floating_ip = 1
    route_map   = "ipv6"
  }
  nat_map_cfg {
    nat_map   = 1
    route_map = "ipv6"
  }
  nat64_cfg {
    nat64     = 1
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
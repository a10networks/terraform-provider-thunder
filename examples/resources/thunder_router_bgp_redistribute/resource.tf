provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_redistribute" "thunder_router_bgp_redistribute" {

  as_number = "300"
  connected_cfg {
    connected = 1
    route_map = "MAP"
  }
  floating_ip_cfg {
    floating_ip = 1
    route_map   = "MAP"
  }
  nat_map_cfg {
    nat_map   = 1
    route_map = "MAP"
  }
  lw4o6_cfg {
    lw4o6     = 1
    route_map = "MAP"
  }
  static_nat_cfg {
    static_nat = 1
    route_map  = "MAP"
  }
  ip_nat_cfg {
    ip_nat    = 1
    route_map = "MAP"
  }
  ip_nat_list_cfg {
    ip_nat_list = 1
    route_map   = "MAP"
  }
  isis_cfg {
    isis      = 1
    route_map = "MAP"
  }
  ospf_cfg {
    ospf      = 1
    route_map = "MAP"
  }
  rip_cfg {
    rip       = 1
    route_map = "MAP"
  }
  static_cfg {
    static    = 1
    route_map = "MAP"
  }
  vip {
    only_flagged_cfg {
      only_flagged = 1
      route_map    = "MAP"
    }
    only_not_flagged_cfg {
      only_not_flagged = 1
      route_map        = "MAP"
    }
  }
}
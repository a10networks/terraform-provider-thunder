provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_ospf_redistribute" "resourceRouterOspfRedistributeTest" {
  process_id = 400
  redist_list {
    type        = "bgp"
    metric_type = 1
    metric      = 40
    route_map   = "MAP"
    tag         = 40
  }

  ip_nat             = 1
  metric_ip_nat      = 40
  metric_type_ip_nat = 1
  route_map_ip_nat   = "NAT"
  tag_ip_nat         = 300
  ip_nat_floating_list {
    ip_nat_prefix              = "8.8.8.8/24"
    ip_nat_floating_ip_forward = "9.9.9.9"
  }
  vip_list {
    type_vip        = "only-flagged"
    metric_type_vip = 1
    metric_vip      = 800
    route_map_vip   = "VIP"
    tag_vip         = 500
  }
  vip_floating_list {
    vip_address             = "5.5.5.5"
    vip_floating_ip_forward = "6.6.6.6"
  }
  ospf_list {
    ospf             = 1
    process_id       = 200
    metric_ospf      = 50
    metric_type_ospf = 1
    route_map_ospf   = "OSPF"
    tag_ospf         = 500
  }
}
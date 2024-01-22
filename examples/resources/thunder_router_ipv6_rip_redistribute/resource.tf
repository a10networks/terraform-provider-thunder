provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_ipv6_rip_redistribute" "thunder_router_ipv6_rip_redistribute" {
  redist_list {
    type      = "bgp"
    metric    = 1
    route_map = "38"
  }
  vip_list {
    vip_type      = "only-flagged"
    vip_metric    = 6
    vip_route_map = "90"
  }
}
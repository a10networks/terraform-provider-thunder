provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_rip_redistribute" "thunder_router_rip_redistribute" {
  redist_list {
    type      = "bgp"
    metric    = 5
    route_map = "74"
  }
  vip_list {
    vip_type      = "only-flagged"
    vip_metric    = 6
    vip_route_map = "33"
  }
}
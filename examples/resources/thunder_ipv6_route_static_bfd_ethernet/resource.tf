provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ipv6_route_static_bfd_ethernet" "eth_1" {
  eth_num         = 2
  nexthop_ipv6_ll = "fe80:1:2:3::4"
  template        = "TEST_2"
  threshold       = 58
  action          = "down"
}

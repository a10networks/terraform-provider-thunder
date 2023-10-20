provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ipv6_route_static_bfd_trunk" "bfd_trunk" {
  trunk_num       = 7
  nexthop_ipv6_ll = "fe80:2:2:8::4"
  template        = "TEST_3"
  threshold       = 240
  action          = "down"
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ipv6_route_static_bfd_bfd_ipv6" "ipv6_1" {
  local_ipv6   = "fda1:2:3:4::1"
  nexthop_ipv6 = "fdb1::1"
  template     = "TEST_1"
  threshold    = 251
  action       = "down"
}
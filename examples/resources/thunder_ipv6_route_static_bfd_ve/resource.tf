provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_route_static_bfd_ve" "VE_1" {
  ve_num          = 55
  nexthop_ipv6_ll = "fe80:55:3:4::1"
  template        = "TEST_4"
  threshold       = 110
  action          = "down"
}
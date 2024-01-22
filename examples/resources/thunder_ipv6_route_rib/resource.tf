provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_route_rib" "rib_1" {
  ipv6_address = "1:2:3:4::/64"
  ipv6_nexthop_ipv6 {
    ipv6_nexthop = "fe80:5:6:7::1"
    ethernet     = 2
    distance     = 71
    description  = "TEST_1_IPV6"
  }
}
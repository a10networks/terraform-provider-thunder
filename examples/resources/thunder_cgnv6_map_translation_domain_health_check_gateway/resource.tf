provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_map_translation_domain_health_check_gateway" "thunder_cgnv6_map_translation_domain_health_check_gateway" {

  name = "test"
  address_list {
    ipv4_gateway = "10.10.10.10"
  }
  ipv6_address_list {
    ipv6_gateway = "2001:db8:3333:4444:5555:6666:7777:8888"
  }
  withdraw_route = "any-link-failure"
}
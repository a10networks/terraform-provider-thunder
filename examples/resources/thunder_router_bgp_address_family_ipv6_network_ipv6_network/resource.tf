provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_address_family_ipv6_network_ipv6_network" "thunderRouterBgpAddressFamilyIpv6NetworkIpv6NetworkTest" {

  as_number    = "300"
  network_ipv6 = "2222::1/64"
  route_map    = "NwRouteMap1"
  backdoor     = 1
  description  = "Ipv6NwIpv6NW1"
  comm_value   = "internet"
}
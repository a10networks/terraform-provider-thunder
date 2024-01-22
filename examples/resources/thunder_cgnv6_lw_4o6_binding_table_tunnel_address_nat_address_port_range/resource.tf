provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lw_4o6_binding_table_tunnel_address_nat_address_port_range" "thunder_cgnv6_lw_4o6_binding_table_tunnel_address_nat_address_port_range" {

  name                    = "test1"
  ipv4_nat_addr           = "10.10.10.10"
  ipv6_tunnel_addr        = "2001:db8:3223:1144:5545:6666:7777:8888"
  port_end                = 60835
  port_start              = 9235
  tunnel_endpoint_address = "2001:db8:3333:4444:5555:6666:7777:8888"
}
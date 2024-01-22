provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lw_4o6_binding_table" "thunder_cgnv6_lw_4o6_binding_table" {
  name = "test"
  tunnel_address_list {
    ipv6_tunnel_addr = "2001:db8:3333:4444:5555:6666:7777:8888"
    user_tag         = "test"
    nat_address_list {
      ipv4_nat_addr = "10.10.10.10"
      user_tag      = "test"
      port_range_list {
        port_start              = 29
        port_end                = 46771
        tunnel_endpoint_address = "2011:db9:3333:4444:5555:6666:7774:8881"
      }
    }
  }
  user_tag = "48"
}
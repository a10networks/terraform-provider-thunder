provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lw_4o6_binding_table_tunnel_address" "thunder_cgnv6_lw_4o6_binding_table_tunnel_address" {

  name             = "test1"
  ipv6_tunnel_addr = "2001:db8:3333:4444:5555:6666:7777:8888"
  nat_address_list {
    ipv4_nat_addr = "10.10.10.10"
    user_tag      = "123"

  }
  user_tag = "80"
}
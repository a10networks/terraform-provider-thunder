provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat46_stateless_static_dest_mapping" "thunder_cgnv6_nat46_stateless_static_dest_mapping" {
  shared     = 1
  v4_address = "10.10.10.10"
  v6_address = "2001:db8:3333:4444:5555:6666:7777:8888"
  vrid       = 0
}
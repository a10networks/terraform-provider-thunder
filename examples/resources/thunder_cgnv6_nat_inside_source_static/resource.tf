provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_inside_source_static" "thunder_cgnv6_nat_inside_source_static" {
  nat_address = "10.10.10.12"
  partition   = "test"
  src_address = "10.10.10.10"
  vrid        = 3
}
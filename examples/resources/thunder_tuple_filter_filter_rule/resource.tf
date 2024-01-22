provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tuple_filter_filter_rule" "thunder_tuple_filter_filter_rule" {

  name     = "test"
  id1      = 2
  src_addr = "10.10.10.10/32"
  dst_addr = "10.10.10.12/32"
  src_port = 20
  dst_port = 22
}
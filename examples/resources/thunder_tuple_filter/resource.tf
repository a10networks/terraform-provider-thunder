provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tuple_filter" "thunder_tuple_filter" {
  name = "test"
  filter_rule_list {
    id1      = 2
    src_addr = "10.10.10.10/32"
    dst_addr = "10.10.10.12/32"
    src_port = 20
    dst_port = 22
  }
  filter_type = "ipv4"
  user_tag    = "testing"
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_endpoint_independent_mapping_tcp" "thunder_cgnv6_lsn_endpoint_independent_mapping_tcp" {
  port_list {
    port     = 100
    port_end = 200
  }
}
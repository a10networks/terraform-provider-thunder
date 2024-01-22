provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_endpoint_independent_filtering_tcp" "thunder_cgnv6_lsn_endpoint_independent_filtering_tcp" {

  port_list {
    port     = 100
    port_end = 1000
  }
  session_limit = 65535
}
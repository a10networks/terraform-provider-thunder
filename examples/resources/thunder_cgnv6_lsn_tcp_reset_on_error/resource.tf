provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_tcp_reset_on_error" "thunder_cgnv6_lsn_tcp_reset_on_error" {
  outbound = "disable"
}
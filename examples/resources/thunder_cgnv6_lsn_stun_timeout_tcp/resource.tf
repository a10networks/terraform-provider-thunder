provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_stun_timeout_tcp" "thunder_cgnv6_lsn_stun_timeout_tcp" {
  port_end   = 20
  port_start = 10
  timeout    = 4
}
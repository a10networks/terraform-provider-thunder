provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_stun_timeout_udp" "thunder_cgnv6_lsn_stun_timeout_udp" {

  port_end   = 20
  port_start = 10
  timeout    = 4

}
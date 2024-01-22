provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_alg_rtp" "thunder_cgnv6_lsn_alg_rtp" {
  rtp_stun_timeout = 3
}
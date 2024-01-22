provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_alg_sip" "thunder_cgnv6_lsn_alg_sip" {
  rtp_stun_timeout = 5
  sampling_enable {
    counters1 = "all"
  }
  sip_value = "enable"
}
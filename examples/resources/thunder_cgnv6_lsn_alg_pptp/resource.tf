provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_alg_pptp" "thunder_cgnv6_lsn_alg_pptp" {
  pptp_value = "enable"
  sampling_enable {
    counters1 = "all"
  }
}
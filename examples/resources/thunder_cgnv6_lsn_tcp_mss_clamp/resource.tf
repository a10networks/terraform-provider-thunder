provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_tcp_mss_clamp" "thunder_cgnv6_lsn_tcp_mss_clamp" {

  min            = 30
  mss_clamp_type = "subtract"
  mss_subtract   = 2

}
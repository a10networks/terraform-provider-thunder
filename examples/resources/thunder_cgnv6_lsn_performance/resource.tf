provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_performance" "thunder_cgnv6_lsn_performance" {
  sampling_enable {
    counters1 = "all"
  }
}
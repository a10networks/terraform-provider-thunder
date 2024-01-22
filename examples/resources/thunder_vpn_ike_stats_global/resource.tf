provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vpn_ike_stats_global" "thunder_vpn_ike_stats_global" {
  sampling_enable {
    counters1 = "all"
  }
}
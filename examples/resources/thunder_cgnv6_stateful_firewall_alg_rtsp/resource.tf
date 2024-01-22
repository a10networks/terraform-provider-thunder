provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_alg_rtsp" "thunder_cgnv6_stateful_firewall_alg_rtsp" {
  rtsp_value = "disable"
  sampling_enable {
    counters1 = "all"
  }
}
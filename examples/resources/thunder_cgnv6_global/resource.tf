provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_global" "thunder_cgnv6_global" {
  ping_sweep_detection = "disable"
  port_scan_detection  = "disable"
  sampling_enable {
    counters1 = "all"
  }
}
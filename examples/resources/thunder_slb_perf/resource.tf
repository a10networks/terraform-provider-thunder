provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_perf" "thunder_slb_perf" {
  sampling_enable {
    counters1 = "all"
  }
}
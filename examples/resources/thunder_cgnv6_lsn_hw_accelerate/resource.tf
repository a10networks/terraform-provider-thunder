provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_hw_accelerate" "thunder_cgnv6_lsn_hw_accelerate" {
  sampling_enable {
    counters1 = "all"
  }
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_one_to_one_global" "thunder_cgnv6_one_to_one_global" {
  mapping_timeout = 10
  sampling_enable {
    counters1 = "all"
  }
}
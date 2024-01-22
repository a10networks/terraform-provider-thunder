provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_pcp" "thunder_cgnv6_pcp" {
  sampling_enable {
    counters1 = "all"
  }
}
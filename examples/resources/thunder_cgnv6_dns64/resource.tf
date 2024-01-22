provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_dns64" "thunder_cgnv6_dns64" {
  sampling_enable {
    counters1 = "all"
  }
}
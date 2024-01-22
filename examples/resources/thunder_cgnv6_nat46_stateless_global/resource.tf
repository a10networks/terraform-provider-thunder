provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat46_stateless_global" "thunder_cgnv6_nat46_stateless_global" {
  sampling_enable {
    counters1 = "all"
  }
}
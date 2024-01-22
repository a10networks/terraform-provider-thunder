provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_fixed_nat_alg_pptp" "thunder_cgnv6_fixed_nat_alg_pptp" {
  sampling_enable {
    counters1 = "all"
  }
}
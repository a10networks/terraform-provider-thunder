provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_fixed_nat_alg_ftp" "thunder_cgnv6_fixed_nat_alg_ftp" {
  sampling_enable {
    counters1 = "all"
  }
}
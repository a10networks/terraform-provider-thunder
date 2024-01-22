provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_histogram" "thunder_cgnv6_nat_histogram" {
  bin_count = 51
  bin_skew  = 71
  data_skew = 21
}
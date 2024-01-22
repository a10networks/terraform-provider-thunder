provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_alg_h323" "thunder_cgnv6_nat64_alg_h323" {
  h323_enable = "enable"
}
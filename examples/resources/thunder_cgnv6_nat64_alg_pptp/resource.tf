provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_alg_pptp" "thunder_cgnv6_nat64_alg_pptp" {
  pptp_enable = "enable"
}
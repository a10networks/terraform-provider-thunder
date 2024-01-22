provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_alg_dns" "thunder_cgnv6_lsn_alg_dns" {
  dns_value = "disable"
}
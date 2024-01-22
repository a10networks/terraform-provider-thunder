provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat46_stateless_prefix" "thunder_cgnv6_nat46_stateless_prefix" {

  ipv6_prefix = "46::/96"
  vrid        = 3

}
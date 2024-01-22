provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_icmpv6" "thunder_cgnv6_nat_icmpv6" {
  respond_to_ping = 1
}
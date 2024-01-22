provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_icmp" "thunder_cgnv6_nat_icmp" {
  always_source_nat_errors = 0
  respond_to_ping          = 0
}
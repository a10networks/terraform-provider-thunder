provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_nat_icmpv6" "natIcmpV6" {
  respond_to_ping = 1
}
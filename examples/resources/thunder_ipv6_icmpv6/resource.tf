provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_icmpv6" "testname" {
  redirect    = 0
  unreachable = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nptv6_common" "thunder_cgnv6_nptv6_common" {
  send_icmpv6_on_error = "disable"
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_sctp_rate_limit_destination" "thunder_cgnv6_sctp_rate_limit_destination" {
  ip         = "10.10.10.10"
  rate_limit = 166196
}
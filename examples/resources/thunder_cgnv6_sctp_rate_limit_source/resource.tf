provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_sctp_rate_limit_source" "thunder_cgnv6_sctp_rate_limit_source" {
  ip         = "10.10.10.10"
  rate_limit = 134522
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_sctp_global" "thunder_cgnv6_sctp_global" {

  half_open_timeout = 5
  idle_timeout      = 6

}
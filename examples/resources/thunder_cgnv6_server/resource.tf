provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_server" "thunder_cgnv6_server" {
  name = "test12345"
  host = "10.10.10.37"

  user_tag = "43"
}
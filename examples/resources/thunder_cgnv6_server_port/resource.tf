provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_server_port" "thunder_cgnv6_server_port" {

  name        = "test12345"
  action      = "enable"
  port_number = 54519
  protocol    = "tcp"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "16"
}
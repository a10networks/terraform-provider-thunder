provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_dns64_virtualserver_port" "thunder_cgnv6_dns64_virtualserver_port" {

  name        = "test"
  port_number = 22866
  protocol    = "dns-udp"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "15"
}
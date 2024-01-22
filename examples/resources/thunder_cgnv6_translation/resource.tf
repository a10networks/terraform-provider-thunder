provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_translation" "thunder_cgnv6_translation" {

  icmp_timeout {
    icmp_timeout_val = 7583
  }
  service_timeout_list {
    service_type = "tcp"
    port         = 11918
    port_end     = 44548
    timeout_val  = 14114
  }
  tcp_timeout = 300
  udp_timeout = 300
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_logging_disable_log_by_destination_ip" "thunder_cgnv6_template_logging_disable_log_by_destination_ip" {

  name      = "test"
  icmp      = 1
  ipv4_addr = "10.10.10.10/32"
  others    = 1
  tcp_list {
    tcp_port_start = 123
    tcp_port_end   = 1234
  }
  udp_list {
    udp_port_start = 123
    udp_port_end   = 1234
  }
  user_tag = "103"
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_logging_disable_log_by_destination" "thunder_cgnv6_template_logging_disable_log_by_destination" {

  name = "test"
  icmp = 0
  ip_list {
    ipv4_addr = "10.10.10.10/32"
    tcp_list {
      tcp_port_start = 124
      tcp_port_end   = 1245
    }
    udp_list {
      udp_port_start = 124
      udp_port_end   = 1245
    }
    icmp     = 1
    others   = 1
    user_tag = "75"
  }

  others = 1
  tcp_list {
    tcp_port_start = 123
    tcp_port_end   = 1234
  }
  udp_list {
    udp_port_start = 123
    udp_port_end   = 1234
  }
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor_disable_log_by_destination" "thunder_netflow_monitor_disable_log_by_destination" {

  name = "a11"
  icmp = 1
  ip_list {
    ipv4_addr = "10.10.10.10/24"
    tcp_list {
      tcp_port_start = 35324
      tcp_port_end   = 35324
    }
    udp_list {
      udp_port_start = 32422
      udp_port_end   = 32422
    }
    icmp     = 1
    others   = 1
    user_tag = "87"
  }
  ip6_list {
    ipv6_addr = "b184:baa0:dff5:52a2:78a2:7c9e:b1f0:4111/24"
    tcp_list {
      tcp_port_start = 31442
      tcp_port_end   = 31442
    }
    udp_list {
      udp_port_start = 23133
      udp_port_end   = 23133
    }
    icmp     = 1
    others   = 1
    user_tag = "19"
  }
  others = 0
  tcp_list {
    tcp_port_start = 43553
    tcp_port_end   = 43553
  }
  udp_list {
    udp_port_start = 32423
    udp_port_end   = 32423
  }
}
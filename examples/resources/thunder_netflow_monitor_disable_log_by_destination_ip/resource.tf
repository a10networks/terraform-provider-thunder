provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor_disable_log_by_destination_ip" "thunder_netflow_monitor_disable_log_by_destination_ip" {

  name      = "a11"
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
  user_tag = "11"
}
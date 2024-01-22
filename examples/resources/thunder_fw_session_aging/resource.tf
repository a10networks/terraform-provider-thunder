provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_session_aging" "thunder_fw_session_aging" {
  name              = "temp"
  icmp_idle_timeout = 300
  ip_idle_timeout   = 33
  user_tag          = "test-user"
  tcp {
    port_cfg {
      tcp_port         = 20
      tcp_idle_timeout = 33
    }

  }
  udp {
    udp_idle_timeout = 222
  }
}
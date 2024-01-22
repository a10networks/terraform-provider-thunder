provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_session_aging_tcp" "thunder_fw_session_aging_tcp" {

  name = "temp"
  port_cfg {
    tcp_port         = 20
    tcp_idle_timeout = 33
  }
}
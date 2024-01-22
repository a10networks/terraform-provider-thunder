provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_server_port" "thunder_fw_server_port" {

  name         = "test"
  action       = "enable"
  health_check = "ping"
  port_number  = 53
  protocol     = "tcp"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "test"
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_log_server_port" "thunder_acos_events_log_server_port" {

  name        = "test"
  port_number = 123
  protocol    = "tcp"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "76"
}
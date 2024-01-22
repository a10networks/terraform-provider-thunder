provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_server_port" "Slb_Server_Port_Test" {

  name              = "test-server1"
  protocol          = "tcp"
  weight            = 10
  stats_data_action = "stats-data-disable"
  template_port     = "default"
  conn_limit        = 23241
  support_http2     = 1
  sampling_enable {
    counters1 = "all"
  }
  no_ssl               = 1
  port_number          = 800
  extended_stats       = 1
  conn_resume          = 2551
  user_tag             = "test123"
  range                = 201
  health_check_disable = 0
  health_check         = "ping"
}
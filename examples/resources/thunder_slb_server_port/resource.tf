provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_server_port" "Slb_Server_Port_Test" {
  server_name       = thunder_server.resourceServerTest.name
  protocol          = "tcp"
  weight            = 10
  stats_data_action = "stats-data-disable"
  template_port     = "incedo"
  conn_limit        = 1
  support_http2     = 1
  sampling_enable {
    counters1 = "all"
  }
  no_ssl              = 1
  template_server_ssl = "a10networks"
  alternate_port {
    alternate_name        = "rs1-a1"
    alternate             = 1
    alternate_server_port = 80
  }
  port_number    = 800
  extended_stats = 1
  conn_resume    = 1
  user_tag       = "test123"
  range          = 200
  auth_cfg {
    service_principal_name = "string1234"
  }
  action       = "disable-with-health-check"
  no_logging   = 1
  health_check = "test"
}
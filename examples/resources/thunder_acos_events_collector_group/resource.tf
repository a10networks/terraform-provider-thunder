provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_collector_group" "thunder_acos_events_collector_group" {

  facility         = "local0"
  format           = "syslog"
  log_distribution = "round-robin"
  log_server_list {
    name = "test"
    port = 346
  }
  name     = "test"
  protocol = "tcp"
  rate     = 500
  sampling_enable {
    counters1 = "all"
  }
  server_distribution_hash = "name"
  use_mgmt_port            = 1
  user_tag                 = "76"
}
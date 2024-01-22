provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vcs_vcs_para" "thunder_vcs_vcs_para" {
  chassis_id                = 8
  config_info               = "93"
  config_seq                = "59"
  dead_interval             = 10
  failure_retry_count_value = 2
  force_wait_interval       = 5
  memory_stat_interval      = 30
  multicast_ip              = "224.0.1.210"
  multicast_port            = 41217
  slog_level                = 7
  slog_method               = 1
  ssl_enable                = 1
  tcp_channel_monitor       = 1
  time_interval             = 3
  transmit_fragment_size    = 6000
}
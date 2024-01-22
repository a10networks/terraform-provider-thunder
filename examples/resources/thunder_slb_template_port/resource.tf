provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_port" "thunder_slb_template_port" {
  name                       = "232"
  bw_rate_limit              = 8388609
  bw_rate_limit_duration     = 128
  bw_rate_limit_resume       = 8388609
  conn_limit                 = 64000000
  conn_rate_limit            = 524290
  dest_nat                   = 1
  down_grace_period          = 4301
  del_session_on_server_down = 1
  dscp                       = 37
  dynamic_member_priority    = 3
  extended_stats             = 1
  health_check               = "ping"
  health_check_disable       = 0
  stats_data_action          = "stats-data-enable"
  user_tag                   = "67"
  weight                     = 1
}
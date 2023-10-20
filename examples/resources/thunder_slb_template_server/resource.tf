provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_server" "server1" {
  name                       = "server1"
  slow_start                 = 1
  initial_slow_start         = 128
  add                        = 3600
  every                      = 20
  till                       = 4000
  bw_rate_limit              = 2
  bw_rate_limit_resume       = 2
  bw_rate_limit_duration     = 1
  bw_rate_limit_no_logging   = 1
  max_dynamic_server         = 3
  min_ttl_ratio              = 3
  health_check               = "ping"
  spoofing_cache             = 1
  bw_rate_limit_acct         = "to-server-only"
  conn_limit                 = 630000
  resume                     = 3
  conn_limit_no_logging      = 1
  conn_rate_limit            = 500
  rate_interval              = "second"
  conn_rate_limit_no_logging = 1
  dns_fail_interval          = 600
  dynamic_server_prefix      = "www"
  dns_query_interval         = 1300
  extended_stats             = 1
  log_selection_failure      = 1
  weight                     = 30
  stats_data_action          = "stats-data-enable"
  user_tag                   = "servertemplate"
}
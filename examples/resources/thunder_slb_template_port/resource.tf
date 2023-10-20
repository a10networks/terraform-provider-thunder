provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_port" "template-port" {
  name                       = "templateport1"
  slow_start                 = 1
  initial_slow_start         = 128
  add                        = 3
  every                      = 5
  till                       = 30000
  conn_limit                 = 2
  resume                     = 1
  conn_limit_no_logging      = 1
  weight                     = 10
  extended_stats             = 1
  del_session_on_server_down = 1
  bw_rate_limit              = 9000
  bw_rate_limit_resume       = 8900
  bw_rate_limit_duration     = 30
  bw_rate_limit_no_logging   = 1
  conn_rate_limit            = 300
  rate_interval              = "100ms"
  conn_rate_limit_no_logging = 0
  dampening_flaps            = 10
  flap_period                = 30
  restore_svc_time           = 20
  dest_nat                   = 1
  down_grace_period          = 300
  dscp                       = 62
  dynamic_member_priority    = 1
  decrement                  = 1
  health_check               = "ping"
  inband_health_check        = 1
  retry                      = 7
  no_ssl                     = 1
  request_rate_limit         = 30000
  request_rate_interval      = "100ms"
  reset                      = 1
  request_rate_no_logging    = 1
  sub_group                  = 5
  source_nat                 = "pool1"
  user_tag                   = "porttemplate1"
}
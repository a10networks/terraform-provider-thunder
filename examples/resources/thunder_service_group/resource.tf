provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_service_group" "service_group_tcp" {
  name                           = "slb_service_tcp"
  protocol                       = "tcp"
  template_port                  = "temp_port"
  template_policy                = "temp_policy"
  lb_method                      = "dst-ip-hash"
  stateless_auto_switch          = 1
  stateless_lb_method2           = "stateless-dst-ip-hash"
  conn_rate                      = 5
  conn_rate_duration             = 10
  conn_rate_grace_period         = 300
  conn_rate_log                  = 1
  min_active_member              = 3
  min_active_member_action       = "dynamic-priority"
  reset_on_server_selection_fail = 1
  priority_affinity              = 1
  reset_priority_affinity        = 1
  backup_server_event_log        = 1
  strict_select                  = 1
  stats_data_action              = "stats-data-disable"
  extended_stats                 = 1
  traffic_replication_mirror     = 1
  health_check                   = "tcp1"
  priorities {
    priority        = 5
    priority_action = "drop"
  }
  sample_rsp_time = 1
  rpt_ext_server  = 1
  report_delay    = 3
  top_slowest     = 0
  top_fastest     = 0
  persist_scoring = "enable"
  user_tag        = "service_group_tcp"
  sampling_enable {
    counters1 = "server_selection_fail_drop"
  }
  member_list {
    name                      = "web1"
    port                      = 80
    member_state              = "disable-with-health-check"
    member_stats_data_disable = 1
    member_template           = "temp_port"
    member_priority           = 3
    user_tag                  = "port"
    sampling_enable {
      counters1 = "total_conn"
    }
  }
}

resource "thunder_service_group" "service_group_udp" {
  name                       = "slb_service_udp"
  protocol                   = "udp"
  template_port              = "temp_port"
  template_policy            = "temp_policy"
  lb_method                  = "dst-ip-hash"
  stateless_auto_switch      = 1
  stateless_lb_method2       = "stateless-dst-ip-hash"
  conn_rate                  = 5
  conn_rate_duration         = 10
  conn_rate_grace_period     = 300
  conn_rate_log              = 1
  min_active_member          = 3
  min_active_member_action   = "dynamic-priority"
  priority_affinity          = 1
  reset_priority_affinity    = 1
  backup_server_event_log    = 1
  strict_select              = 1
  stats_data_action          = "stats-data-disable"
  extended_stats             = 1
  traffic_replication_mirror = 1
  health_check               = "tcp1"
  priorities {
    priority        = 5
    priority_action = "drop"
  }
  persist_scoring = "enable"
  user_tag        = "service_group_udp"
  sampling_enable {
    counters1 = "server_selection_fail_drop"
  }
  member_list {
    name                      = "web1"
    port                      = 53
    member_state              = "disable-with-health-check"
    member_stats_data_disable = 1
    member_template           = "temp_port"
    member_priority           = 3
    user_tag                  = "port"
    sampling_enable {
      counters1 = "total_conn"
    }
  }
}
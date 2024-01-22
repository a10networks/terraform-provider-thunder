provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_service_group" "thunder_slb_service_group" {

  name                    = "sg90"
  protocol                = "tcp"
  backup_server_event_log = 1
  extended_stats          = 1
  health_check            = "ping"
  member_list {
    name                      = "test-server1"
    port                      = 22556
    member_state              = "enable"
    member_stats_data_disable = 1
    member_template           = "default"
    member_priority           = 1
    user_tag                  = "114"
    sampling_enable {
      counters1 = "all"
    }
  }
  min_active_member        = 897
  min_active_member_action = "dynamic-priority"
  persist_scoring          = "global"
  priorities {
    priority        = 13
    priority_action = "proceed"
  }
  priority_affinity              = 1
  reset_on_server_selection_fail = 1
  reset_priority_affinity        = 1
  sample_rsp_time                = 1
  sampling_enable {
    counters1 = "all"
  }
  stateless_lb_method        = "stateless-dst-ip-hash"
  stats_data_action          = "stats-data-enable"
  traffic_replication_mirror = 1
  user_tag                   = "22"
}
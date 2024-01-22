provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_logging" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_logging" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    log_dropped           = 1
    conn_tcp_dropped      = 1
  }
}
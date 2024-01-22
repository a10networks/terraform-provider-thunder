provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_tcp_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_tcp_trigger_stats_rate" {

  name                  = "test"
  attemptfails          = 1
  duration              = 60
  noroute               = 1
  threshold_exceeded_by = 5
}
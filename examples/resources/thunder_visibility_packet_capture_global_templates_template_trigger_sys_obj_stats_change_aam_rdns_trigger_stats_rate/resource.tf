provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_rdns_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_rdns_trigger_stats_rate" {

  name                  = "test"
  duration              = 60
  request_dropped       = 1
  response_error        = 1
  response_failure      = 1
  response_timeout      = 1
  threshold_exceeded_by = 5
}
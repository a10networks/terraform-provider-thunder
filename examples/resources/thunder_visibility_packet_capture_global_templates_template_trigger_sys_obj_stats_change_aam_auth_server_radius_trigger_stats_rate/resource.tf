provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_radius_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_radius_trigger_stats_rate" {

  name                  = "test"
  accounting_failure    = 1
  authen_failure        = 1
  authorize_failure     = 1
  duration              = 60
  job_start_error       = 1
  other_error           = 1
  polling_control_error = 1
  request_dropped       = 1
  response_error        = 1
  response_failure      = 1
  response_timeout      = 1
  threshold_exceeded_by = 5
  timeout_error         = 1
}
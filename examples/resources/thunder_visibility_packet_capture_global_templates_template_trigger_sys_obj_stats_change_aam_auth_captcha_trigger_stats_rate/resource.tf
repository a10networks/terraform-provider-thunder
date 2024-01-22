provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_captcha_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_captcha_trigger_stats_rate" {

  name                  = "test"
  attr_fail             = 1
  duration              = 60
  job_start_error       = 1
  json_fail             = 1
  other_error           = 1
  polling_control_error = 1
  request_dropped       = 1
  response_error        = 1
  response_failure      = 1
  response_timeout      = 1
  threshold_exceeded_by = 5
  timeout_error         = 1
}
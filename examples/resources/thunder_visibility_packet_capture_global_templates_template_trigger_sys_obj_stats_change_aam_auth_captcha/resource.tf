provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_captcha" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_captcha" {

  name = "test"
  trigger_stats_inc {
    request_dropped       = 1
    response_failure      = 1
    response_error        = 1
    response_timeout      = 1
    json_fail             = 1
    attr_fail             = 1
    timeout_error         = 1
    other_error           = 1
    job_start_error       = 1
    polling_control_error = 1
  }
}
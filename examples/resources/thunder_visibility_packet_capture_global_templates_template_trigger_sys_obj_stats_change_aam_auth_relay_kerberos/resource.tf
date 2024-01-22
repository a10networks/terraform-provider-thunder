provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_relay_kerberos" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_relay_kerberos" {

  name = "test"
  trigger_stats_inc {
    timeout_error         = 1
    other_error           = 1
    request_dropped       = 1
    response_failure      = 1
    response_error        = 1
    response_timeout      = 1
    job_start_error       = 1
    polling_control_error = 1
  }
}
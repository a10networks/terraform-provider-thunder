provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_ldap_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_ldap_trigger_stats_rate" {

  name                  = "test"
  admin_bind_failure    = 1
  authorize_failure     = 1
  bind_failure          = 1
  duration              = 60
  job_start_error       = 1
  other_error           = 1
  polling_control_error = 1
  pw_change_failure     = 1
  request_dropped       = 1
  response_error        = 1
  response_failure      = 1
  response_timeout      = 1
  search_failure        = 1
  ssl_session_failure   = 1
  threshold_exceeded_by = 5
  timeout_error         = 1
}
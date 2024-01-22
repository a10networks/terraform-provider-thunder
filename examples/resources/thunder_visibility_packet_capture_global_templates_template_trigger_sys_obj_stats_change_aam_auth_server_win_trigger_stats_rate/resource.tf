provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_win_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_win_trigger_stats_rate" {

  name                                 = "test"
  duration                             = 60
  kerberos_delete_kdc_keytab_failure   = 1
  kerberos_generate_kdc_keytab_failure = 1
  kerberos_job_start_error             = 1
  kerberos_other_error                 = 1
  kerberos_polling_control_error       = 1
  kerberos_pw_change_failure           = 1
  kerberos_pw_expiry                   = 1
  kerberos_request_dropped             = 1
  kerberos_response_error              = 1
  kerberos_response_failure            = 1
  kerberos_response_timeout            = 1
  kerberos_timeout_error               = 1
  kerberos_validate_kdc_failure        = 1
  ntlm_authentication_failure          = 1
  ntlm_job_start_error                 = 1
  ntlm_other_error                     = 1
  ntlm_polling_control_error           = 1
  ntlm_prepare_req_failed              = 1
  ntlm_proto_negotiation_failure       = 1
  ntlm_request_dropped                 = 1
  ntlm_response_error                  = 1
  ntlm_response_failure                = 1
  ntlm_response_timeout                = 1
  ntlm_session_setup_failed            = 1
  ntlm_timeout_error                   = 1
  threshold_exceeded_by                = 5
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl" "thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl" {

  name           = "test"
  capture_config = "test"
  trigger_stats_rate {
    threshold_exceeded_by      = 5
    duration                   = 60
    krb_timeout_error          = 1
    krb_other_error            = 1
    krb_pw_expiry              = 1
    krb_pw_change_failure      = 1
    ntlm_proto_nego_failure    = 1
    ntlm_session_setup_failure = 1
    ntlm_prepare_req_error     = 1
    ntlm_auth_failure          = 1
    ntlm_timeout_error         = 1
    ntlm_other_error           = 1
    krb_validate_kdc_failure   = 1
  }
  trigger_stats_severity {
    error          = 1
    error_alert    = 1
    error_warning  = 1
    error_critical = 1
    drop           = 1
    drop_alert     = 1
    drop_warning   = 1
    drop_critical  = 1
  }
  user_tag = "14"
}
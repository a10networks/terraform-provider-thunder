provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl_trigger_stats_inc" {

  name                       = "test"
  krb_other_error            = 1
  krb_pw_change_failure      = 1
  krb_pw_expiry              = 1
  krb_timeout_error          = 1
  krb_validate_kdc_failure   = 1
  ntlm_auth_failure          = 1
  ntlm_other_error           = 1
  ntlm_prepare_req_error     = 1
  ntlm_proto_nego_failure    = 1
  ntlm_session_setup_failure = 1
  ntlm_timeout_error         = 1
}
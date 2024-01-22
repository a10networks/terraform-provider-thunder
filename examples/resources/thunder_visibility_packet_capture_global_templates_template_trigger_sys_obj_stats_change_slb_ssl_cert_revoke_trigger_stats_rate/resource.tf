provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ssl_cert_revoke_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ssl_cert_revoke_trigger_stats_rate" {

  name                         = "test"
  crl_cache_status_revoked     = 1
  crl_connection_error         = 1
  crl_other_error              = 1
  crl_response_status_revoked  = 1
  crl_response_status_unknown  = 1
  crl_uri_https                = 1
  crl_uri_not_found            = 1
  crl_uri_unsupported          = 1
  duration                     = 60
  ocsp_cache_miss              = 1
  ocsp_cache_status_revoked    = 1
  ocsp_chain_status_revoked    = 1
  ocsp_chain_status_unknown    = 1
  ocsp_connection_error        = 1
  ocsp_other_error             = 1
  ocsp_response_no_nonce       = 1
  ocsp_response_nonce_error    = 1
  ocsp_response_status_revoked = 1
  ocsp_response_status_unknown = 1
  ocsp_uri_https               = 1
  ocsp_uri_not_found           = 1
  ocsp_uri_unsupported         = 1
  threshold_exceeded_by        = 5
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ssl_forward_proxy" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ssl_forward_proxy" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by              = 5
    duration                           = 60
    failed_in_ssl_handshakes           = 1
    failed_in_crypto_operations        = 1
    failed_in_tcp                      = 1
    failed_in_certificate_verification = 1
    failed_in_certificate_signing      = 1
    invalid_ocsp_stapling_response     = 1
    revoked_ocsp_response              = 1
    unsupported_ssl_version            = 1
    connections_failed                 = 1
  }
}
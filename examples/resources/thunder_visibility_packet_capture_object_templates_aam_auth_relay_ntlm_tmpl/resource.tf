provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl" {

  name           = "test"
  capture_config = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    failure               = 1
    buffer_alloc_fail     = 1
    encoding_fail         = 1
    insert_header_fail    = 1
    parse_header_fail     = 1
    internal_error        = 1
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
  user_tag = "35"
}
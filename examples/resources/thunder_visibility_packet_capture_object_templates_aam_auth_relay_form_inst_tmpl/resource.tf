provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_form_inst_tmpl" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_form_inst_tmpl" {

  name           = "test"
  capture_config = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    invalid_srv_rsp       = 1
    post_fail             = 1
    invalid_cred          = 1
    bad_req               = 1
    not_fnd               = 1
    error                 = 1
    other_error           = 1
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
  user_tag = "101"
}
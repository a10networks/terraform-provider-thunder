provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_captcha_inst_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_auth_captcha_inst_tmpl_trigger_stats_rate" {

  name                  = "test"
  attr_fail             = 1
  duration              = 60
  json_fail             = 1
  other_error           = 1
  parse_fail            = 1
  threshold_exceeded_by = 5
  timeout_error         = 1
}
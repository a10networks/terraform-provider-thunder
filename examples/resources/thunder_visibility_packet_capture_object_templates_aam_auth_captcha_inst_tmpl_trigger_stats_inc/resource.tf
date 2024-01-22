provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_captcha_inst_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_auth_captcha_inst_tmpl_trigger_stats_inc" {

  name          = "test"
  attr_fail     = 1
  json_fail     = 1
  other_error   = 1
  parse_fail    = 1
  timeout_error = 1
}
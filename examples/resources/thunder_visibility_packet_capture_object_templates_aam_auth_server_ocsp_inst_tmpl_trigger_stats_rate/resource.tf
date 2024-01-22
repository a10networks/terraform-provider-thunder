provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_server_ocsp_inst_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_auth_server_ocsp_inst_tmpl_trigger_stats_rate" {

  name                  = "test"
  duration              = 60
  fail                  = 1
  stapling_fail         = 1
  stapling_timeout      = 1
  threshold_exceeded_by = 5
  timeout               = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ws_fed_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ws_fed_tmpl_trigger_stats_rate" {

  name                  = "test"
  duration              = 60
  failure               = 1
  threshold_exceeded_by = 5
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ws_fed_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ws_fed_tmpl_trigger_stats_inc" {

  name    = "test"
  failure = 1
}
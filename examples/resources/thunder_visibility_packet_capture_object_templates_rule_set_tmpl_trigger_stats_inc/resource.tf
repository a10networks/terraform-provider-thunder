provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_rule_set_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_rule_set_tmpl_trigger_stats_inc" {

  name            = "test"
  deny            = 1
  reset           = 1
  unmatched_drops = 1
}
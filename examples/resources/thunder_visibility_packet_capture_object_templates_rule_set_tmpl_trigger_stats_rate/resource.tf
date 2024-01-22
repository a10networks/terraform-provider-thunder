provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_rule_set_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_rule_set_tmpl_trigger_stats_rate" {

  name                  = "test"
  deny                  = 1
  duration              = 60
  reset                 = 1
  threshold_exceeded_by = 5
  unmatched_drops       = 1
}
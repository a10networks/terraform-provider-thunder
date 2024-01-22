provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_cgnv6_dns64_vs_port_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_cgnv6_dns64_vs_port_tmpl_trigger_stats_rate" {

  name                     = "test"
  duration                 = 60
  es_total_failure_actions = 1
  threshold_exceeded_by    = 5
}
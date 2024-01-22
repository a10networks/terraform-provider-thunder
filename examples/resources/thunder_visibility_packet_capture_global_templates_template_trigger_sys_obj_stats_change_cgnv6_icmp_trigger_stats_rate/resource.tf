provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_icmp_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_icmp_trigger_stats_rate" {

  name                  = "test"
  duration              = 60
  icmp_to_icmp_err      = 1
  icmp_to_icmpv6_err    = 1
  icmpv6_to_icmp_err    = 1
  icmpv6_to_icmpv6_err  = 1
  threshold_exceeded_by = 5
}
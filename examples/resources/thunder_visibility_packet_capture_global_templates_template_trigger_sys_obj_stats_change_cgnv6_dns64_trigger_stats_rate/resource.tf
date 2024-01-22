provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_rate" {

  name                  = "test"
  drop                  = 1
  duration              = 60
  query_bad_pkt         = 1
  resp_bad_pkt          = 1
  resp_bad_qr           = 1
  threshold_exceeded_by = 5
}
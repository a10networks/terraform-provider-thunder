provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_inc" {

  name          = "test"
  drop          = 1
  query_bad_pkt = 1
  resp_bad_pkt  = 1
  resp_bad_qr   = 1
}
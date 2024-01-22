provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64" {

  name = "test"
  trigger_stats_inc {
    query_bad_pkt = 1
    resp_bad_pkt  = 1
    resp_bad_qr   = 1
    drop          = 1
  }
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_sip_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_sip_trigger_stats_rate" {

  name                  = "test"
  duration              = 60
  method_unknown        = 1
  parse_error           = 1
  tcp_out_of_order_drop = 1
  threshold_exceeded_by = 5
}
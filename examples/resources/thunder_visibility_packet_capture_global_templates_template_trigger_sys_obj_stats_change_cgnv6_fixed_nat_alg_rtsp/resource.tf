provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_alg_rtsp" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_alg_rtsp" {

  name = "test"
  trigger_stats_inc {
    stream_creation_failure = 1
    port_allocation_failure = 1
    no_session_mem          = 1
  }
}
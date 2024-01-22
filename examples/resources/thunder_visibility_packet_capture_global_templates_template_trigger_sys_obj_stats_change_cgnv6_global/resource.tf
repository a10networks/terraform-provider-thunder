provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_global" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_global" {

  name = "test"
  trigger_stats_inc {
    udp_total_ports_allocated  = 1
    icmp_total_ports_allocated = 1
  }
}
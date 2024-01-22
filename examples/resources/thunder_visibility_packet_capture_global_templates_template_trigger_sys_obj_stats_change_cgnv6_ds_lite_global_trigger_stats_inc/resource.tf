provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global_trigger_stats_inc" {

  name                      = "test"
  fullcone_failure          = 1
  nat_port_unavailable_icmp = 1
  nat_port_unavailable_tcp  = 1
  nat_port_unavailable_udp  = 1
  user_quota_failure        = 1
}
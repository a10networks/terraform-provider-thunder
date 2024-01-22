provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global" {

  name = "test"
  trigger_stats_inc {
    user_quota_failure        = 1
    nat_port_unavailable_tcp  = 1
    nat_port_unavailable_udp  = 1
    nat_port_unavailable_icmp = 1
    fullcone_failure          = 1
  }
}
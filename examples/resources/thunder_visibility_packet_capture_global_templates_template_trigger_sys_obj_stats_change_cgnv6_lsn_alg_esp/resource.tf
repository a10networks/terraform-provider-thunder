provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_esp" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_esp" {

  name = "test"
  trigger_stats_inc {
    nat_ip_conflict = 1
  }
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
  }
}
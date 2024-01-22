provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_ctr_lib_acct" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_ctr_lib_acct" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by     = 5
    duration                  = 60
    total_nodes_free_failed   = 1
    total_nodes_unlink_failed = 1
  }
}
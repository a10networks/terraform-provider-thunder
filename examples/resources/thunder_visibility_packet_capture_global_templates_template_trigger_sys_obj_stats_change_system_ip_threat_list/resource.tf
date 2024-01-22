provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_ip_threat_list" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_ip_threat_list" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by    = 5
    duration                 = 60
    error_out_of_memory      = 1
    error_out_of_spe_entries = 1
  }
}
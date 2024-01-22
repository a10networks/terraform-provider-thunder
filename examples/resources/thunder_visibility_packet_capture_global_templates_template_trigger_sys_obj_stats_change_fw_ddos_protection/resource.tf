provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_ddos_protection" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_ddos_protection" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by              = 5
    duration                           = 60
    ddos_entries_too_many              = 1
    ddos_entry_add_to_bgp_failure      = 1
    ddos_entry_remove_from_bgp_failure = 1
    ddos_packet_dropped                = 1
  }
}
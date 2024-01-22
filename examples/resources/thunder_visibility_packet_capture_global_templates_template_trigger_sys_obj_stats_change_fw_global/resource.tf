provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by     = 5
    duration                  = 60
    fullcone_creation_failure = 1
  }
}
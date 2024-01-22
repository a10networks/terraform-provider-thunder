provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_alg_rtsp" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_alg_rtsp" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by   = 5
    duration                = 60
    transport_alloc_failure = 1
  }
}
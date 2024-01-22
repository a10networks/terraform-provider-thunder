provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_tcp_syn_cookie_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_tcp_syn_cookie_trigger_stats_rate" {

  name                  = "test"
  duration              = 60
  threshold_exceeded_by = 5
  verification_failed   = 1
}
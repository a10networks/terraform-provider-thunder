provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_service_group_mem_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_auth_service_group_mem_tmpl_trigger_stats_rate" {

  name                  = "test"
  curr_conn_overflow    = 1
  duration              = 60
  threshold_exceeded_by = 5
}
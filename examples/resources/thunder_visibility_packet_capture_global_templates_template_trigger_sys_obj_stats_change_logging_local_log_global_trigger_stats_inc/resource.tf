provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_logging_local_log_global_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_logging_local_log_global_trigger_stats_inc" {

  name          = "test"
  enqueue_error = 1
  enqueue_full  = 1
}
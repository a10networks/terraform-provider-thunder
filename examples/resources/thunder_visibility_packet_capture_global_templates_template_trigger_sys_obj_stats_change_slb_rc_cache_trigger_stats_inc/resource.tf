provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache_trigger_stats_inc" {

  name                  = "test"
  content_toobig        = 1
  content_toosmall      = 1
  entry_create_failures = 1
  rv_failure            = 1
}
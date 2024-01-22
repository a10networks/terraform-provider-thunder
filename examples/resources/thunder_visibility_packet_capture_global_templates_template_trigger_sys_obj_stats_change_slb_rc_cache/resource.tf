provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    rv_failure            = 1
    content_toobig        = 1
    content_toosmall      = 1
    entry_create_failures = 1
  }
}
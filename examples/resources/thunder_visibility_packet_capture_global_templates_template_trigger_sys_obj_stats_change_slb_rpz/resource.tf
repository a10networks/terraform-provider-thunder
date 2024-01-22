provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rpz" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rpz" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    set_bw_error          = 1
    parse_error           = 1
  }
}
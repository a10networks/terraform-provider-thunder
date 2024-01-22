provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_crl_srcip_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_crl_srcip_trigger_stats_inc" {

  name              = "test"
  out_of_sessions   = 1
  threshold_exceed  = 1
  too_many_sessions = 1
}
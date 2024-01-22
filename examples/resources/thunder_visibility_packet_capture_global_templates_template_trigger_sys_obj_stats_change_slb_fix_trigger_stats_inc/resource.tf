provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_fix_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_fix_trigger_stats_inc" {

  name        = "test"
  client_err  = 1
  noroute     = 1
  server_err  = 1
  snat_fail   = 1
  svrsel_fail = 1
}
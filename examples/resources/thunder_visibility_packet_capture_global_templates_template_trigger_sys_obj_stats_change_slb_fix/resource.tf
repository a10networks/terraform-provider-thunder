provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_fix" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_fix" {

  name = "test"
  trigger_stats_inc {
    svrsel_fail = 1
    noroute     = 1
    snat_fail   = 1
    client_err  = 1
    server_err  = 1
  }
}
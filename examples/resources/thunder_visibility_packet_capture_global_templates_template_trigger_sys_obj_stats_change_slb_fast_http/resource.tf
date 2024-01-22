provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_fast_http" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_fast_http" {

  name = "test"
  trigger_stats_inc {
    parsereq_fail       = 1
    svrsel_fail         = 1
    fwdreq_fail         = 1
    fwdreqdata_fail     = 1
    snat_fail           = 1
    full_proxy_fpga_err = 1
    req_over_limit      = 1
    req_rate_over_limit = 1
  }
}
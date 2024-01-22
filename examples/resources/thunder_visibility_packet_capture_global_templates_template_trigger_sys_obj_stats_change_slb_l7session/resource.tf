provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l7session" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l7session" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    conn_not_exist        = 1
    wbuf_cb_failed        = 1
    err_event             = 1
    err_cb_failed         = 1
    server_conn_failed    = 1
    server_select_fail    = 1
    data_cb_failed        = 1
    hps_fwdreq_fail       = 1
  }
}
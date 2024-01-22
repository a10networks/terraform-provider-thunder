provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_icap_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_icap_trigger_stats_rate" {

  name                     = "test"
  app_serv_conn_err        = 1
  app_serv_conn_no_pcb_err = 1
  chunk1_hdr_err           = 1
  chunk2_hdr_err           = 1
  chunk_bad_trail_err      = 1
  duration                 = 60
  encap_hdr_incomplete_err = 1
  http_resp_hdr_err        = 1
  http_resp_line_parse_err = 1
  http_resp_line_read_err  = 1
  icap_line_err            = 1
  icap_ver_err             = 1
  no_icap_resp_err         = 1
  no_payload_buff_err      = 1
  no_payload_next_buff_err = 1
  no_status_code_err       = 1
  prep_req_fail_err        = 1
  req_hdr_incomplete_err   = 1
  resp_hdr_err             = 1
  resp_hdr_incomplete_err  = 1
  resp_line_parse_err      = 1
  resp_line_read_err       = 1
  serv_sel_fail_err        = 1
  start_icap_conn_fail_err = 1
  threshold_exceeded_by    = 5
}
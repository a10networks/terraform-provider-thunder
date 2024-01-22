provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_smtp_vport_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_smtp_vport_tmpl_trigger_stats_inc" {

  name                          = "test"
  forward_req_data_fail         = 1
  forward_req_fail              = 1
  get_all_headers_fail          = 1
  insert_resonse_line_fail      = 1
  line_extend_fail              = 1
  line_table_extend_fail        = 1
  line_too_long                 = 1
  no_proxy                      = 1
  parse_req_fail                = 1
  parse_request_line_fail       = 1
  parse_resonse_line_fail       = 1
  read_request_line_fail        = 1
  recv_server_unknow_reply_code = 1
  remove_resonse_line_fail      = 1
  send_client_service_not_ready = 1
  server_select_fail            = 1
  server_starttls_fail          = 1
  snat_fail                     = 1
  too_many_headers              = 1
}
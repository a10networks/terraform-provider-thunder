provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_spdy_proxy_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_spdy_proxy_trigger_stats_inc" {

  name                        = "test"
  compress_ctx_alloc_fail     = 1
  ctx_alloc_fail              = 1
  data_cb_no_tuple            = 1
  data_no_stream              = 1
  data_no_stream_goaway_close = 1
  data_no_stream_no_goaway    = 1
  decompress_fail             = 1
  est_cb_no_tuple             = 1
  header_compress_fail        = 1
  http_conn_alloc_fail        = 1
  http_err_stream_closed      = 1
  invalid_frame_size          = 1
  invalid_version             = 1
  request_header_alloc_fail   = 1
  session_err                 = 1
  stream_alloc_fail           = 1
  stream_err                  = 1
  stream_not_found            = 1
  tcp_err                     = 1
}
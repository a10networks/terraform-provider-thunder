provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_http2" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_http2" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by           = 5
    duration                        = 60
    protocol_error                  = 1
    internal_error                  = 1
    proxy_alloc_error               = 1
    split_buff_fail                 = 1
    invalid_frame_size              = 1
    error_max_invalid_stream        = 1
    data_no_stream                  = 1
    flow_control_error              = 1
    settings_timeout                = 1
    frame_size_error                = 1
    refused_stream                  = 1
    cancel                          = 1
    compression_error               = 1
    connect_error                   = 1
    enhance_your_calm               = 1
    inadequate_security             = 1
    http_1_1_required               = 1
    deflate_alloc_fail              = 1
    inflate_alloc_fail              = 1
    inflate_header_fail             = 1
    bad_connection_preface          = 1
    cant_allocate_control_frame     = 1
    cant_allocate_settings_frame    = 1
    bad_frame_type_for_stream_state = 1
    wrong_stream_state              = 1
    data_queue_alloc_error          = 1
    buff_alloc_error                = 1
    cant_allocate_rst_frame         = 1
    cant_allocate_goaway_frame      = 1
    cant_allocate_ping_frame        = 1
    cant_allocate_stream            = 1
  }
}
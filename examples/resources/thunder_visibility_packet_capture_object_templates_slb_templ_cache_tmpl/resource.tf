provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_slb_templ_cache_tmpl" "thunder_visibility_packet_capture_object_templates_slb_templ_cache_tmpl" {

  name           = "test"
  capture_config = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    nc_req_header         = 1
    nc_res_header         = 1
    rv_failure            = 1
    content_toobig        = 1
    content_toosmall      = 1
    entry_create_failures = 1
    header_save_error     = 1
  }
  trigger_stats_severity {
    error          = 1
    error_alert    = 1
    error_warning  = 1
    error_critical = 1
    drop           = 1
    drop_alert     = 1
    drop_warning   = 1
    drop_critical  = 1
  }
  user_tag = "106"
}
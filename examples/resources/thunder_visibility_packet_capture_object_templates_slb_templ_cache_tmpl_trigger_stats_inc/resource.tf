provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_slb_templ_cache_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_slb_templ_cache_tmpl_trigger_stats_inc" {

  name                  = "test"
  content_toobig        = 1
  content_toosmall      = 1
  entry_create_failures = 1
  header_save_error     = 1
  nc_req_header         = 1
  nc_res_header         = 1
  rv_failure            = 1
}
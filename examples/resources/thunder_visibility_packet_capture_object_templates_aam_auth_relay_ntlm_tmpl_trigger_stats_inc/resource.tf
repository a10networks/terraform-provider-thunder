provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl_trigger_stats_inc" {

  name               = "test"
  buffer_alloc_fail  = 1
  encoding_fail      = 1
  failure            = 1
  insert_header_fail = 1
  internal_error     = 1
  parse_header_fail  = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_form_inst_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_form_inst_tmpl_trigger_stats_inc" {

  name            = "test"
  bad_req         = 1
  error           = 1
  invalid_cred    = 1
  invalid_srv_rsp = 1
  not_fnd         = 1
  other_error     = 1
  post_fail       = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_generic_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_generic_trigger_stats_inc" {

  name                      = "test"
  client_fail               = 1
  client_select_fail        = 1
  dcmsg_error               = 1
  invalid_avp               = 1
  mismatch_fwd_id           = 1
  mismatch_rev_id           = 1
  no_fwd_tuple              = 1
  no_rev_tuple              = 1
  no_route                  = 1
  no_session_id             = 1
  reply_error_info_fail     = 1
  reply_unknown_session_id  = 1
  retry_client_request_fail = 1
  server_fail               = 1
  snat_fail                 = 1
  svrsel_fail               = 1
  unkwn_cmd_code            = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_pptp_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_pptp_trigger_stats_rate" {

  name = "test"

  call_reply_pns_call_id_mismatch = 1
  call_req_pns_call_id_mismatch   = 1
  duration                        = 60
  no_gre_session_match            = 1
  threshold_exceeded_by           = 5
}
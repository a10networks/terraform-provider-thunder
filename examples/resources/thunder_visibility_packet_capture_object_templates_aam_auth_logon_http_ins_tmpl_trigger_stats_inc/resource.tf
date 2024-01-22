provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_logon_http_ins_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_auth_logon_http_ins_tmpl_trigger_stats_inc" {

  name           = "test"
  spn_krb_faiure = 1
}
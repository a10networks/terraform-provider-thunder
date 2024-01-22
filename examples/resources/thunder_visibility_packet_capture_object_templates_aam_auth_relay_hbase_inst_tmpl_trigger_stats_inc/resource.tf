provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_inc" {

  name         = "test"
  bad_req      = 1
  forbidden    = 1
  no_creds     = 1
  not_found    = 1
  server_error = 1
  unauth       = 1
  unavailable  = 1
}
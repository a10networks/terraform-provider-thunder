provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_rate" {

  name                  = "test"
  bad_req               = 1
  duration              = 60
  forbidden             = 1
  no_creds              = 1
  not_found             = 1
  server_error          = 1
  threshold_exceeded_by = 5
  unauth                = 1
  unavailable           = 1
}
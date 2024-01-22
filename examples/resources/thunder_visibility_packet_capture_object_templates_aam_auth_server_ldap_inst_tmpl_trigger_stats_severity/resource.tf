provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_server_ldap_inst_tmpl_trigger_stats_severity" "thunder_visibility_packet_capture_object_templates_aam_auth_server_ldap_inst_tmpl_trigger_stats_severity" {

  name           = "test"
  drop           = 1
  drop_alert     = 1
  drop_critical  = 1
  drop_warning   = 1
  error          = 1
  error_alert    = 1
  error_critical = 1
  error_warning  = 1
}
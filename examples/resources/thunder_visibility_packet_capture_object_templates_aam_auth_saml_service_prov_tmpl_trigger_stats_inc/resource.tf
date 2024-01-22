provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_saml_service_prov_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_auth_saml_service_prov_tmpl_trigger_stats_inc" {

  name           = "test"
  acs_authz_fail = 1
  acs_error      = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_saml_id_prov_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_auth_saml_id_prov_tmpl_trigger_stats_rate" {

  name                  = "test"
  acs_fail              = 1
  duration              = 60
  md_fail               = 1
  threshold_exceeded_by = 5
}
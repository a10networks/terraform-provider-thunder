provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_saml_global" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_saml_global" {

  name = "test"
  trigger_stats_inc {
    acs_authz_fail = 1
    acs_error      = 1
  }
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_jwt_authorization_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_jwt_authorization_tmpl_trigger_stats_inc" {

  name                  = "test"
  jwt_authorize_failure = 1
  jwt_missing_claim     = 1
  jwt_missing_token     = 1
  jwt_other_error       = 1
  jwt_signature_failure = 1
  jwt_token_expired     = 1
}
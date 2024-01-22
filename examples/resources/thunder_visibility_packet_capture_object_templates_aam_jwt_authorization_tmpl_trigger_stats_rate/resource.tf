provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_jwt_authorization_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_jwt_authorization_tmpl_trigger_stats_rate" {

  name                  = "test"
  duration              = 60
  jwt_authorize_failure = 1
  jwt_missing_claim     = 1
  jwt_missing_token     = 1
  jwt_other_error       = 1
  jwt_signature_failure = 1
  jwt_token_expired     = 1
  threshold_exceeded_by = 5
}
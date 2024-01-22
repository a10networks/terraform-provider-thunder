provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_validation_policy" "thunder_template_gtp_validation_policy" {
  name                     = "test_validation"
  anomaly_action           = "drop"
  anomaly_checks           = "enable"
  anti_spoofing_action     = "drop"
  anti_spoofing_check      = "disable"
  crosslayer_corr_action   = "drop"
  crosslayer_correlation   = "disable"
  mandatory_ie_action      = "drop"
  mandatory_ie_check       = "enable"
  msisdn_imsi_corr_action  = "drop"
  msisdn_imsi_correlation  = "disable"
  out_of_order_ie_action   = "drop"
  out_of_order_ie_check    = "disable"
  out_of_state_ie_action   = "drop"
  out_of_state_ie_check    = "disable"
  reserved_ie_action       = "drop"
  reserved_ie_check        = "enable"
  sequence_num_corr_action = "drop"
  sequence_num_correlation = "disable"
  user_tag                 = "29"
}
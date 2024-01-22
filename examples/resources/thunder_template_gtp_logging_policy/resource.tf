provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_logging_policy" "thunder_template_gtp_logging_policy" {
  name = "test"
  log {
    mandatory_ie_check       = 1
    out_of_state_ie_check    = 1
    out_of_order_ie_check    = 1
    invalid_teid_check       = 1
    reserved_ie_check        = 1
    message_filtering        = 1
    apn_imsi_filtering       = 1
    rat_type_filtering       = 1
    msisdn_filtering         = 1
    crosslayer_correlation   = 1
    anti_spoofing_check      = 1
    msisdn_imsi_correlation  = 1
    max_message_length_check = 1
    gtp_in_gtp_filtering     = 1
    sequence_num_correlation = 1
    invalid_header_check     = 1
  }
  user_tag = "85"
}
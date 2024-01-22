provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_message_filtering_policy_version_v1" "thunder_template_gtp_message_filtering_policy_version_v1" {

  name                  = "test"
  create_mbms           = "disable"
  create_pdp            = "enable"
  delete_mbms           = "disable"
  delete_pdp            = "enable"
  enable_disable_action = "enable"
  gtp_pdu               = "enable"
  initiate_pdp          = "enable"
  mbms_deregistration   = "disable"
  mbms_notification     = "disable"
  mbms_registration     = "disable"
  mbms_session          = "disable"
  ms_info_change        = "enable"
  pdu_notification      = "enable"
  reserved_messages     = "disable"
  update_mbms           = "disable"
  update_pdp            = "enable"
}
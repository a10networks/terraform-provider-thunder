provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_message_filtering_policy_version_v0" "thunder_template_gtp_message_filtering_policy_version_v0" {

  name                  = "test"
  create_aa_pdp         = "enable"
  create_pdp            = "enable"
  delete_aa_pdp         = "enable"
  delete_pdp            = "enable"
  enable_disable_action = "enable"
  gtp_pdu               = "enable"
  pdu_notification      = "enable"
  reserved_messages     = "disable"
  update_pdp            = "enable"
}
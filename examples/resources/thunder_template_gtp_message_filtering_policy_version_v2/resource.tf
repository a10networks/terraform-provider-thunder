provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_message_filtering_policy_version_v2" "thunder_template_gtp_message_filtering_policy_version_v2" {

  name                  = "test"
  bearer_resource       = "enable"
  change_notification   = "enable"
  create_bearer         = "enable"
  create_session        = "enable"
  delete_bearer         = "enable"
  delete_command        = "enable"
  delete_pdn            = "enable"
  delete_session        = "enable"
  enable_disable_action = "enable"
  modify_bearer         = "enable"
  modify_command        = "enable"
  pgw_downlink_trigger  = "disable"
  remote_ue_report      = "enable"
  reserved_messages     = "disable"
  resume                = "enable"
  suspend               = "enable"
  trace_session         = "disable"
  update_bearer         = "enable"
  update_pdn            = "enable"
}
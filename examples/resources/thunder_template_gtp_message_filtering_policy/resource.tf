provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_message_filtering_policy" "thunder_template_gtp_message_filtering_policy" {
  name           = "test"
  interface_type = "roaming"
  user_tag       = "119"
  version_v0 {
    enable_disable_action = "enable"
    create_pdp            = "enable"
    update_pdp            = "enable"
    delete_pdp            = "enable"
    pdu_notification      = "enable"
    gtp_pdu               = "enable"
    create_aa_pdp         = "enable"
    delete_aa_pdp         = "enable"
    reserved_messages     = "disable"
  }
  version_v1 {
    enable_disable_action = "enable"
    create_pdp            = "enable"
    update_pdp            = "enable"
    delete_pdp            = "enable"
    initiate_pdp          = "enable"
    pdu_notification      = "enable"
    ms_info_change        = "enable"
    gtp_pdu               = "enable"
    mbms_session          = "disable"
    mbms_notification     = "disable"
    mbms_registration     = "disable"
    mbms_deregistration   = "disable"
    create_mbms           = "disable"
    delete_mbms           = "disable"
    update_mbms           = "disable"
    reserved_messages     = "disable"
  }
  version_v2 {
    enable_disable_action = "enable"
    change_notification   = "enable"
    create_session        = "enable"
    delete_session        = "enable"
    modify_bearer         = "enable"
    remote_ue_report      = "enable"
    modify_command        = "enable"
    delete_command        = "enable"
    bearer_resource       = "enable"
    create_bearer         = "enable"
    update_bearer         = "enable"
    delete_bearer         = "enable"
    delete_pdn            = "enable"
    update_pdn            = "enable"
    suspend               = "enable"
    resume                = "enable"
    pgw_downlink_trigger  = "disable"
    trace_session         = "disable"
    reserved_messages     = "disable"
  }
}
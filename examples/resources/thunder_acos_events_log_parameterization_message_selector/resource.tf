provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_log_parameterization_message_selector" "thunder_acos_events_log_parameterization_message_selector" {
  rule_list {
    index         = 167
    action        = "send"
    severity_oper = "equal-and-higher"
    severity_val  = "emergency"
    user_tag      = "15"
  }
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_log_parameterization" "thunder_acos_events_log_parameterization" {
  log_rate = 10
  message_selector {
    rule_list {
      index         = 88
      action        = "drop"
      severity_oper = "equal"
      severity_val  = "alert"
      user_tag      = "9"
    }
  }
}
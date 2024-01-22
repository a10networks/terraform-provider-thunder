provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_message_selector_rule" "thunder_acos_events_message_selector_rule" {

  name          = "test"
  action        = "send"
  index         = 26
  severity_oper = "equal-and-higher"
  severity_val  = "emergency"
  user_tag      = "116"
}
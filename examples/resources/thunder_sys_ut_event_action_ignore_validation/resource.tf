provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sys_ut_event_action_ignore_validation" "thunder_sys_ut_event_action_ignore_validation" {

  event_number = 18
  direction    = "expect"
  all          = 1
  l1           = 0
  l2           = 0
  l3           = 0
  l4           = 0
}
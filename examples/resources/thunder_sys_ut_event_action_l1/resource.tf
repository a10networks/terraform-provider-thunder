provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sys_ut_event_action_l1" "thunder_sys_ut_event_action_l1" {

  event_number = 18
  direction    = "expect"
  auto         = 0
  length       = 1
  value        = 537
}
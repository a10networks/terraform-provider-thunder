provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sys_ut_event_action" "thunder_sys_ut_event_action" {

  event_number = 18
  direction    = "expect"
  drop         = 1
}
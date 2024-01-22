provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sys_ut_event_action_l2" "thunder_sys_ut_event_action_l2" {

  event_number = 18
  direction    = "expect"
  ethertype    = 1
  value        = 5326
}
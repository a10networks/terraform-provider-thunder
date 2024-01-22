provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sys_ut_event" "thunder_sys_ut_event" {
  event_number = 18
  user_tag     = "67"
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sys_ut_event_action_l3" "thunder_sys_ut_event_action_l3" {

  event_number = 18
  direction    = "expect"
  ip_list {
    src_dst      = "dest"
    ipv4_address = "10.2.5.6"
  }
  protocol = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ospf_ifsm" "thunder_debug_ospf_ifsm" {
  events = 0
  status = 0
  timers = 0
}
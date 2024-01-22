provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ospf_nfsm" "thunder_debug_ospf_nfsm" {
  events = 0
  status = 0
  timers = 0
}
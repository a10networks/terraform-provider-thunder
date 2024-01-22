provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ospf_events" "thunder_debug_ospf_events" {
  abr    = 0
  asbr   = 0
  os     = 0
  router = 0
  vlink  = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ipv6_rip" "thunder_debug_ipv6_rip" {
  all    = 0
  detail = 0
  events = 0
  nsm    = 0
  packet = 0
  recv   = 0
  send   = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_sctp" "thunder_debug_sctp" {
  level  = 1
  packet = 1
}
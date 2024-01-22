provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_scaleout" "thunder_debug_scaleout" {
  config       = 0
  debug_level  = 3
  event        = 0
  packet       = 0
  session_sync = 0
}
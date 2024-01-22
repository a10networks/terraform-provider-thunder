provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_lacp" "thunder_debug_lacp" {
  all    = 0
  cli    = 0
  detail = 0
  event  = 0
  ha     = 0
  packet = 0
  sync   = 0
  timer  = 0
}
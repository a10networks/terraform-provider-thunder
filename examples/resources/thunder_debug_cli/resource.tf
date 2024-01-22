provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_cli" "thunder_debug_cli" {
  all    = 0
  io     = 0
  parser = 0
}
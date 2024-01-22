provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_logging" "thunder_debug_logging" {
  all     = 0
  command = 0
  error   = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_backtrace" "thunder_debug_backtrace" {
  dumy = 0
}
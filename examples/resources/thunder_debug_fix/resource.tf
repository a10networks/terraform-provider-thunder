provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_fix" "thunder_debug_fix" {
  dummy = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_management" "thunder_debug_management" {
  all    = 0
  system = 0
}
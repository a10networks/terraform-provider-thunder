provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_l4" "thunder_debug_l4" {
  dumy = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_mgcp" "thunder_debug_mgcp" {
  dumy = 0
}
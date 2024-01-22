provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_imish" "thunder_debug_imish" {
  dumy = 0
}
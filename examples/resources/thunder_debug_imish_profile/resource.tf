provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_imish_profile" "thunder_debug_imish_profile" {
  dumy = 0
}
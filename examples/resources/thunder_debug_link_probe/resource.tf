provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_link_probe" "thunder_debug_link_probe" {
  dumy = 0
}
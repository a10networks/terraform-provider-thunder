provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_polltech_enable" "thunder_debug_polltech_enable" {
  dumy = 0
}
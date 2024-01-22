provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ssl" "thunder_debug_ssl" {
  dumy = 0
}
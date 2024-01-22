provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_tunnel" "thunder_debug_tunnel" {
  dumy = 0
}
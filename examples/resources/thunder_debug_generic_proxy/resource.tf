provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_generic_proxy" "thunder_debug_generic_proxy" {
  debug_level = 1
}
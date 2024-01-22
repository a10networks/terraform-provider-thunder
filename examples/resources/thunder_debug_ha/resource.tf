provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ha" "thunder_debug_ha" {
  debug_level = 2
}
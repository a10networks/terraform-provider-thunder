provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_diameter" "thunder_debug_diameter" {
  dumy = 0
}
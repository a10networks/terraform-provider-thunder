provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_vtep_error" "thunder_debug_vtep_error" {
  dumy = 0
}
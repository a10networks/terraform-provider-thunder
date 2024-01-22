provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_vtep_event" "thunder_debug_vtep_event" {
  dumy = 0
}
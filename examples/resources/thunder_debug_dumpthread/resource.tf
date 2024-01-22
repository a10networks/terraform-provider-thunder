provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_dumpthread" "thunder_debug_dumpthread" {
  dumy = 0
}
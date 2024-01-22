provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_migration" "thunder_debug_migration" {
  dumy = 0
}
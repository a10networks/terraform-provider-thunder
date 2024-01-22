provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_system" "thunder_debug_system" {
  debug_system_enum = "all"
}
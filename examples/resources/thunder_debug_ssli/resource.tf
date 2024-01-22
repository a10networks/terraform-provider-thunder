provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ssli" "thunder_debug_ssli" {
  dumy = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_mlb" "thunder_debug_mlb" {
  dummy = 0
}
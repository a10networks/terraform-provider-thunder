provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_gtp" "thunder_debug_gtp" {
  level  = 1
  packet = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_es" "thunder_debug_es" {
  level = 2
}
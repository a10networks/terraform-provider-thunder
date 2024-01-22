provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_aflow" "thunder_debug_aflow" {
  level = 2
}
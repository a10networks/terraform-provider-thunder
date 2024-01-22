provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_smtp" "thunder_debug_smtp" {
  level = 3
}
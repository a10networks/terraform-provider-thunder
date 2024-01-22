provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_sess" "thunder_debug_sess" {
  dumy = 0
}
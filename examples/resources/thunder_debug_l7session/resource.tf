provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_l7session" "thunder_debug_l7session" {
  level = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_tcp" "thunder_debug_tcp" {
  level = 3
  stack = 1
}
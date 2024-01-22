provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_cache" "thunder_debug_cache" {
  level = 2
}
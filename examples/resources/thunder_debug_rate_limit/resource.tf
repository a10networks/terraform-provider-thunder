provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_rate_limit" "thunder_debug_rate_limit" {
  dumy = 0
}
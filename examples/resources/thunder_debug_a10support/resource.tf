provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_a10support" "thunder_debug_a10support" {
  duration = 2
  password = "20"
}
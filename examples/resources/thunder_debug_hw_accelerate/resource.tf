provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_hw_accelerate" "thunder_debug_hw_accelerate" {
  dummy = 0
}
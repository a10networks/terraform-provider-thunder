provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_fw" "thunder_debug_fw" {
  ddos = 0
}
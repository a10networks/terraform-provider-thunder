provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ip" "thunder_debug_ip" {
  dumy = 0
}
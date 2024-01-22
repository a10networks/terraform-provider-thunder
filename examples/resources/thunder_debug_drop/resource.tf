provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_drop" "thunder_debug_drop" {
  dumy = 0
}
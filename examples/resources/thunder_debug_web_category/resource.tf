provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_web_category" "thunder_debug_web_category" {
  dumy = 0
}
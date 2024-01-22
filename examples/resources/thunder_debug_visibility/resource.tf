provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_visibility" "thunder_debug_visibility" {
  xflow = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_l2_redirect" "thunder_debug_l2_redirect" {
  level = 2
}
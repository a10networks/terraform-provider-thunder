provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_l7_http3" "thunder_debug_l7_http3" {
  level = 5
}
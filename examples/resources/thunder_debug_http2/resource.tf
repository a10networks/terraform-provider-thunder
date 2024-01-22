provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_http2" "thunder_debug_http2" {
  level = 1
}
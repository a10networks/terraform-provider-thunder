provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_http_proxy" "thunder_debug_http_proxy" {
  level = 3
}
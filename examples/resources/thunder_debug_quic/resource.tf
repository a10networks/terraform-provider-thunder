provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_quic" "thunder_debug_quic" {
  dummy     = 1
  level     = 1
  tls       = 1
  tls_level = 2
}
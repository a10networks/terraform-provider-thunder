provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_cgn" "thunder_debug_cgn" {
  drops = 0
  error = 0
}
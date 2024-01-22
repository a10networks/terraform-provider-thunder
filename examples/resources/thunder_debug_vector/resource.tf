provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_vector" "thunder_debug_vector" {
  error  = 0
  packet = 0
  trace  = 0
}
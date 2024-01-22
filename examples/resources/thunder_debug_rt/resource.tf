provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_rt" "thunder_debug_rt" {
  all     = 0
  check   = 0
  command = 0
}
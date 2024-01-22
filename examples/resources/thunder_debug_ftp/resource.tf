provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ftp" "thunder_debug_ftp" {
  dumy = 0
}
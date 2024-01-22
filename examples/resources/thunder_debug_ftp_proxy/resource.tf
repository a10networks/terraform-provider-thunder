provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ftp_proxy" "thunder_debug_ftp_proxy" {
  dummy = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_axdebug_exit" "thunder_axdebug_exit" {
  stop_capture = 0
}
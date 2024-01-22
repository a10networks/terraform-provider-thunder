provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_msg_proxy" "thunder_debug_msg_proxy" {
  dummy = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_running_config" "thunder_running_config" {
  aflex      = 0
  class_list = 0
}
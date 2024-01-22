provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_chassis_info" "thunder_chassis_info" {
  brief = 0
}
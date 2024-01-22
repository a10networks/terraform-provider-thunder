provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_mqtt" "thunder_debug_mqtt" {
  level = 2
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_lro" "thunder_system_lro" {
  disable = 0
}
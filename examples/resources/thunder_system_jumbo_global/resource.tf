provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_jumbo_global" "thunder_system_jumbo_global" {
  enable_jumbo = 0
}
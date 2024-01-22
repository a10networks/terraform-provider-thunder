provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_link_monitor" "thunder_system_link_monitor" {
  disable = 0
}
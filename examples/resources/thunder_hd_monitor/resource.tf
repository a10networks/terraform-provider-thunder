provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_hd_monitor" "thunder_hd_monitor" {
  enable = 0
}
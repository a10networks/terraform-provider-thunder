provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_multi_config" "thunder_multi_config" {
  enable = 0
}
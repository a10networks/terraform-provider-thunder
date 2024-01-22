provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_authentication_console" "thunder_authentication_console" {
  type_cfg {
    type = 0
  }
}
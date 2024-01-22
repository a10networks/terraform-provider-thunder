provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_authorization" "thunder_authorization" {
  debug = 14
  method {
    tacplus = 0
  }
}
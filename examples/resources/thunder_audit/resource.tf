provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_audit" "thunder_audit" {
  enable    = 1
  privilege = 0
  size      = 11075
}
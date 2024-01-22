provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_enable_password" "thunder_enable_password" {
  encrypted = "98"
  password  = "56"
}
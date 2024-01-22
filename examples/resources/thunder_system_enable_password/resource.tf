provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_enable_password" "thunder_system_enable_password" {
  follow_password_policy = 0
}
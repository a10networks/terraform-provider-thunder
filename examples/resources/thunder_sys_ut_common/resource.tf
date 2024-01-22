provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sys_ut_common" "thunder_sys_ut_common" {
  delay            = 3112
  proceed_on_error = 0
}
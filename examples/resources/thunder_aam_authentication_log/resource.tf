provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_log" "thunder_aam_authentication_log" {
  enable   = 0
  facility = "local0"
  format   = "syslog"
}
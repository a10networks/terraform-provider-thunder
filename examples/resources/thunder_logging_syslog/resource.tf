provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_syslog" "thunder_logging_syslog" {
  syslog_levelname = "alert"
}
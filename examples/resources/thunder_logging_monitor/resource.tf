provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_monitor" "thunder_logging_monitor" {
  monitor_levelname = "emergency"
}
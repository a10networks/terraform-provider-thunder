provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_cli_monitor_interval" "thunder_system_cli_monitor_interval" {
  interval = 3858
}
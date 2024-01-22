provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_harmony_controller_telemetry" "thunder_harmony_controller_telemetry" {
  log_rate = 10
}
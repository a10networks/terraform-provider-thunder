provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_reporting_telemetry_export_interval" "thunder_visibility_reporting_telemetry_export_interval" {
  value = 0
}
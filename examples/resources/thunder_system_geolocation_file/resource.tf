provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_geolocation_file" "thunder_system_geolocation_file" {
  error_info {
  }
}
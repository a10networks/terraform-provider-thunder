provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_external_create" "thunder_health_external_create" {
  description = "2"
  file_name   = "23"
}
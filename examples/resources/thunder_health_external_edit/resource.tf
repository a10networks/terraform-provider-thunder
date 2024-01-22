provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_external_edit" "thunder_health_external_edit" {
  description = "56"
  file_name   = "2"
}
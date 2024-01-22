provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_external_delete" "thunder_health_external_delete" {
  file_name = "3"
}
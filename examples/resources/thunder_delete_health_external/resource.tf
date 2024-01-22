provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_delete_health_external" "thunder_delete_health_external" {
  file_name = "14"
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_disable_management_service_http" "thunder_disable_management_service_http" {
  management = 0
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_disable_management_service_https" "thunder_disable_management_service_https" {
  management = 0
}
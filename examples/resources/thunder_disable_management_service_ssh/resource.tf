provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_disable_management_service_ssh" "thunder_disable_management_service_ssh" {
  management = 0
}
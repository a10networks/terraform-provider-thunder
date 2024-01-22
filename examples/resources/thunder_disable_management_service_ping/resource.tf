provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_disable_management_service_ping" "thunder_disable_management_service_ping" {
  all_data_intf = 1
  management    = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_ipmi_service" "thunder_system_ipmi_service" {
  disable = 0
}
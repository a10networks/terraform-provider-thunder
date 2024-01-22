provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_app_protocol_port_monitor" "thunder_ip_app_protocol_port_monitor" {
  disable = 0
  enable  = 1
}
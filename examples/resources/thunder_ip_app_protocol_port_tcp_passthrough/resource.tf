provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_app_protocol_port_tcp_passthrough" "thunder_ip_app_protocol_port_tcp_passthrough" {
  disable = 0
  enable  = 1
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_twamp_responder" "thunder_network_twamp_responder" {
  enable_ip = 1
  port      = 59650
}
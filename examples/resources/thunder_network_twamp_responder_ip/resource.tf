provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_twamp_responder_ip" "thunder_network_twamp_responder_ip" {

  acl_id = 100
}
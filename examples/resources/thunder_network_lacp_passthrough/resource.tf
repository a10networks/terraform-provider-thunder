provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_lacp_passthrough" "thunder_network_lacp_passthrough" {
  peer_from = 1
  peer_to   = 2
}
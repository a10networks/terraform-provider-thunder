provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_arp_timeout" "thunder_network_arp_timeout" {
  timeout = 200
}
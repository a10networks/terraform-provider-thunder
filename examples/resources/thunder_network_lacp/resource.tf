provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_lacp" "thunder_network_lacp" {
  system_priority = 32762
}
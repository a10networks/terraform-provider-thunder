provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_mac_age_time" "thunder_network_mac_age_time" {
  aging_time = 100
}
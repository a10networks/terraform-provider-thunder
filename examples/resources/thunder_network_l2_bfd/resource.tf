provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_l2_bfd" "thunder_network_l2_bfd" {
  ether_type  = "3"
  multiplier  = 4
  rx_interval = 800
  tx_interval = 800
}
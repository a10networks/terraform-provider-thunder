provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_interface_ethernet" "thunder_vrrp_a_interface_ethernet" {
  both         = 1
  ethernet_val = 2
  no_heartbeat = 1
  user_tag     = "test"
}
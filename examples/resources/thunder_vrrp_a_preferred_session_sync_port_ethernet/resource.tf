provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_preferred_session_sync_port_ethernet" "thunder_vrrp_a_preferred_session_sync_port_ethernet" {
  pre_eth  = 1
  pre_vlan = 514
}
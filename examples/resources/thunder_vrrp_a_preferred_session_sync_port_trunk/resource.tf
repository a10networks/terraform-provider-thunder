provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_preferred_session_sync_port_trunk" "thunder_vrrp_a_preferred_session_sync_port_trunk" {

  pre_trunk = 3
  pre_vlan  = 3
}
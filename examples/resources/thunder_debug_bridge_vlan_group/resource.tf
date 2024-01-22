provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_bridge_vlan_group" "thunder_debug_bridge_vlan_group" {
  dumy = 0
}
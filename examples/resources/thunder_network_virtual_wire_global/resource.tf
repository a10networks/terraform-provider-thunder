provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_virtual_wire_global" "thunder_network_virtual_wire_global" {
  sampling_enable {
    counters1 = "all"
  }
  src_mac_learning   = 0
  update_active_vlan = "all"
  vlan_update_period = 195
}
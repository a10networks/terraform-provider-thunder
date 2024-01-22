provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_spanning_tree_mode_rstp" "thunder_network_spanning_tree_mode_rstp" {

  action     = 0
  priority   = 32768
  vlan_start = 1378

}
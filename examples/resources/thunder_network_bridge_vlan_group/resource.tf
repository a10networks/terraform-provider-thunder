provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_bridge_vlan_group" "thunder_network_bridge_vlan_group" {

  bridge_vlan_group_number = 134
  forward_traffic          = "forward-ip-traffic"
  user_tag                 = "13"
  vlan_list {
    vlan_start = 2323
    vlan_end   = 2323
  }
}
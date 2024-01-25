provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_vlan_stats" "thunder_network_vlan_stats" {

  vlan_num = 2730
}
output "get_network_vlan_stats" {
  value = ["${data.thunder_network_vlan_stats.thunder_network_vlan_stats}"]
}
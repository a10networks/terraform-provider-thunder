provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_vlan_oper" "thunder_network_vlan_oper" {

  vlan_num = 2730

}
output "get_network_vlan_oper" {
  value = ["${data.thunder_network_vlan_oper.thunder_network_vlan_oper}"]
}
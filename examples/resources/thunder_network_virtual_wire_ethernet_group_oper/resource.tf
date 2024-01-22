provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_virtual_wire_ethernet_group_oper" "thunder_network_virtual_wire_ethernet_group_oper" {

  group_id = 1

}
output "get_network_virtual_wire_ethernet_group_oper" {
  value = ["${data.thunder_network_virtual_wire_ethernet_group_oper.thunder_network_virtual_wire_ethernet_group_oper}"]
}
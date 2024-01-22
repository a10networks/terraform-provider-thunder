provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_virtual_wire_global_oper" "thunder_network_virtual_wire_global_oper" {
}
output "get_network_virtual_wire_global_oper" {
  value = ["${data.thunder_network_virtual_wire_global_oper.thunder_network_virtual_wire_global_oper}"]
}
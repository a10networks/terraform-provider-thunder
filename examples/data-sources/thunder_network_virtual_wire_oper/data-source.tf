provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_virtual_wire_oper" "thunder_network_virtual_wire_oper" {

  virtual_wire_id = 30
}
output "get_network_virtual_wire_oper" {
  value = ["${data.thunder_network_virtual_wire_oper.thunder_network_virtual_wire_oper}"]
}
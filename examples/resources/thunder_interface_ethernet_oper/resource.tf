provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_ethernet_oper" "thunder_interface_ethernet_oper" {
  ifnum = 2
}
output "get_interface_ethernet_oper" {
  value = ["${data.thunder_interface_ethernet_oper.thunder_interface_ethernet_oper}"]
}
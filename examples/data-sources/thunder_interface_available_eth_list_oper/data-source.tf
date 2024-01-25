provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_available_eth_list_oper" "thunder_interface_available_eth_list_oper" {
}
output "get_interface_available_eth_list_oper" {
  value = ["${data.thunder_interface_available_eth_list_oper.thunder_interface_available_eth_list_oper}"]
}
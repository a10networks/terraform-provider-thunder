provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_mac_address_dynamic_oper" "thunder_network_mac_address_dynamic_oper" {
}
output "get_network_mac_address_dynamic_oper" {
  value = ["${data.thunder_network_mac_address_dynamic_oper.thunder_network_mac_address_dynamic_oper}"]
}
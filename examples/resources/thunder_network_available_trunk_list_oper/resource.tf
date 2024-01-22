provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_available_trunk_list_oper" "thunder_network_available_trunk_list_oper" {

}
output "get_network_available_trunk_list_oper" {
  value = ["${data.thunder_network_available_trunk_list_oper.thunder_network_available_trunk_list_oper}"]
}
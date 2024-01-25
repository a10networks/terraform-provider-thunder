provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_spanning_tree_mode_info_oper" "thunder_network_spanning_tree_mode_info_oper" {
}
output "get_network_spanning_tree_mode_info_oper" {
  value = ["${data.thunder_network_spanning_tree_mode_info_oper.thunder_network_spanning_tree_mode_info_oper}"]
}
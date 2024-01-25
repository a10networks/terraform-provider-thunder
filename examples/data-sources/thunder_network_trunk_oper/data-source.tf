provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_trunk_oper" "thunder_network_trunk_oper" {

}
output "get_network_trunk_oper" {
  value = ["${data.thunder_network_trunk_oper.thunder_network_trunk_oper}"]
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_trunk_oper" "thunder_interface_trunk_oper" {

  ifnum = 11

}
output "get_interface_trunk_oper" {
  value = ["${data.thunder_interface_trunk_oper.thunder_interface_trunk_oper}"]
}
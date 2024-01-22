provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_management_oper" "thunder_interface_management_oper" {
}
output "get_interface_management_oper" {
  value = ["${data.thunder_interface_management_oper.thunder_interface_management_oper}"]
}
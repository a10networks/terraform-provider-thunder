provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_common_oper" "thunder_interface_common_oper" {
}
output "get_interface_common_oper" {
  value = ["${data.thunder_interface_common_oper.thunder_interface_common_oper}"]
}
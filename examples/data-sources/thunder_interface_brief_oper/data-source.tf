provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_brief_oper" "thunder_interface_brief_oper" {
}
output "get_interface_brief_oper" {
  value = ["${data.thunder_interface_brief_oper.thunder_interface_brief_oper}"]
}
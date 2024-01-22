provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_loopback_oper" "thunder_interface_loopback_oper" {

  ifnum = 4
}
output "get_interface_loopback_oper" {
  value = ["${data.thunder_interface_loopback_oper.thunder_interface_loopback_oper}"]
}
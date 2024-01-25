provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_automatic_update_info_oper" "thunder_automatic_update_info_oper" {
}
output "get_automatic_update_info_oper" {
  value = ["${data.thunder_automatic_update_info_oper.thunder_automatic_update_info_oper}"]
}
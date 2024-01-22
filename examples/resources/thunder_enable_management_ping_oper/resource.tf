provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_enable_management_ping_oper" "thunder_enable_management_ping_oper" {
}
output "get_enable_management_ping_oper" {
  value = ["${data.thunder_enable_management_ping_oper.thunder_enable_management_ping_oper}"]
}
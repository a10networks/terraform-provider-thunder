provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_radius_server_oper" "thunder_system_radius_server_oper" {
}
output "get_system_radius_server_oper" {
  value = ["${data.thunder_system_radius_server_oper.thunder_system_radius_server_oper}"]
}
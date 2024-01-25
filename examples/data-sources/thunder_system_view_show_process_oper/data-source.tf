provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_view_show_process_oper" "thunder_system_view_show_process_oper" {
}
output "get_system_view_show_process_oper" {
  value = ["${data.thunder_system_view_show_process_oper.thunder_system_view_show_process_oper}"]
}
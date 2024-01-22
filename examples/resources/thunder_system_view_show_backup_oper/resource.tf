provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_view_show_backup_oper" "thunder_system_view_show_backup_oper" {
}
output "get_system_view_show_backup_oper" {
  value = ["${data.thunder_system_view_show_backup_oper.thunder_system_view_show_backup_oper}"]
}
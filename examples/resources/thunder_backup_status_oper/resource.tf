provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_backup_status_oper" "thunder_backup_status_oper" {
}
output "get_backup_status_oper" {
  value = ["${data.thunder_backup_status_oper.thunder_backup_status_oper}"]
}
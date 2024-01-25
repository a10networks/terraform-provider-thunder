provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_config_sync_status_oper" "thunder_config_sync_status_oper" {
}
output "get_config_sync_status_oper" {
  value = ["${data.thunder_config_sync_status_oper.thunder_config_sync_status_oper}"]
}
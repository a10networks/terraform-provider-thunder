provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_sync_oper" "thunder_ddos_sync_oper" {
}
output "get_ddos_sync_oper" {
  value = ["${data.thunder_ddos_sync_oper.thunder_ddos_sync_oper}"]
}
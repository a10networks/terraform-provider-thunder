provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_sys_audit_log_oper" "thunder_sys_audit_log_oper" {
}
output "get_sys_audit_log_oper" {
  value = ["${data.thunder_sys_audit_log_oper.thunder_sys_audit_log_oper}"]
}
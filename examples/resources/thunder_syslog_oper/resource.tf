provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_syslog_oper" "thunder_syslog_oper" {
}
output "get_syslog_oper" {
  value = ["${data.thunder_syslog_oper.thunder_syslog_oper}"]
}
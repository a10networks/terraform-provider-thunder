provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_logging_local_log_access_log_oper" "thunder_logging_local_log_access_log_oper" {
}
output "get_logging_local_log_access_log_oper" {
  value = ["${data.thunder_logging_local_log_access_log_oper.thunder_logging_local_log_access_log_oper}"]
}
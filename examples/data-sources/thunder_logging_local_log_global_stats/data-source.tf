provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_logging_local_log_global_stats" "thunder_logging_local_log_global_stats" {
}
output "get_logging_local_log_global_stats" {
  value = ["${data.thunder_logging_local_log_global_stats.thunder_logging_local_log_global_stats}"]
}
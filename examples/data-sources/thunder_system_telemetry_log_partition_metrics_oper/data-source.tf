provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_telemetry_log_partition_metrics_oper" "thunder_system_telemetry_log_partition_metrics_oper" {
}
output "get_system_telemetry_log_partition_metrics_oper" {
  value = ["${data.thunder_system_telemetry_log_partition_metrics_oper.thunder_system_telemetry_log_partition_metrics_oper}"]
}
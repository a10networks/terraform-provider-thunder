provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_telemetry_log_top_k_source_list_oper" "thunder_system_telemetry_log_top_k_source_list_oper" {
}
output "get_system_telemetry_log_top_k_source_list_oper" {
  value = ["${data.thunder_system_telemetry_log_top_k_source_list_oper.thunder_system_telemetry_log_top_k_source_list_oper}"]
}
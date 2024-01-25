provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_cpu_data_cpu_overall_oper" "thunder_system_cpu_data_cpu_overall_oper" {
}
output "get_system_cpu_data_cpu_overall_oper" {
  value = ["${data.thunder_system_cpu_data_cpu_overall_oper.thunder_system_cpu_data_cpu_overall_oper}"]
}
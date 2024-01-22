provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_cpu_data_cpu_oper" "thunder_system_cpu_data_cpu_oper" {
}
output "get_system_cpu_data_cpu_oper" {
  value = ["${data.thunder_system_cpu_data_cpu_oper.thunder_system_cpu_data_cpu_oper}"]
}
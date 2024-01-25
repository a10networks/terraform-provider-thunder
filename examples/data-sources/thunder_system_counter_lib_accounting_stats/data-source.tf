provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_counter_lib_accounting_stats" "thunder_system_counter_lib_accounting_stats" {
}
output "get_system_counter_lib_accounting_stats" {
  value = ["${data.thunder_system_counter_lib_accounting_stats.thunder_system_counter_lib_accounting_stats}"]
}
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_sessions_smp_table_oper" "thunder_sessions_smp_table_oper" {
}
output "get_sessions_smp_table_oper" {
  value = ["${data.thunder_sessions_smp_table_oper.thunder_sessions_smp_table_oper}"]
}
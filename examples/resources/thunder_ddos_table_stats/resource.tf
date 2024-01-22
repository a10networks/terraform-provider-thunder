provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_table_stats" "thunder_ddos_table_stats" {
}
output "get_ddos_table_stats" {
  value = ["${data.thunder_ddos_table_stats.thunder_ddos_table_stats}"]
}
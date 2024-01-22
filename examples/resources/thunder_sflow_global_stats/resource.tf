provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_sflow_global_stats" "thunder_sflow_global_stats" {
}
output "get_sflow_global_stats" {
  value = ["${data.thunder_sflow_global_stats.thunder_sflow_global_stats}"]
}